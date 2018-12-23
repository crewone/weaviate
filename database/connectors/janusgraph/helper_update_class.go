package janusgraph

import (
	"fmt"
	"time"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/go-openapi/strfmt"
)

// map properties in thing.Schema according to the mapping.
type edge struct {
	PropertyName string
	Type         string
	Reference    string
	Location     string
}

type edgeFromRefProp struct {
	localEdges   []edge
	networkEdges []edge
	edgesToDrop  []string
}

func (j *Janusgraph) updateClass(k kind.Kind, className schema.ClassName, UUID strfmt.UUID, atContext string, lastUpdateTimeUnix int64, rawProperties interface{}) error {
	vertexLabel := j.state.getMappedClassName(className)

	q := gremlin.G.V().
		HasString(PROP_KIND, k.Name()).
		HasString(PROP_UUID, UUID.String()).
		As("class").
		StringProperty(PROP_CLASS_ID, string(vertexLabel)).
		StringProperty(PROP_AT_CONTEXT, atContext).
		Int64Property(PROP_LAST_UPDATE_TIME_UNIX, lastUpdateTimeUnix)

	q, err := j.addEdgesToQuery(q, k, className, rawProperties)
	if err != nil {
		return err
	}

	_, err = j.client.Execute(q)
	return err
}

func (j *Janusgraph) addEdgesToQuery(q *gremlin.Query, k kind.Kind, className schema.ClassName, rawProperties interface{}) (*gremlin.Query, error) {

	var localEdges []edge
	var networkEdges []edge
	var dropTheseEdgeTypes []string

	properties, ok := rawProperties.(map[string]interface{})
	if !ok {
		// nothing to do because we don't have any
		// (useable) properties
		return q, nil
	}

	for propName, value := range properties {
		sanitizedPropertyName := schema.AssertValidPropertyName(propName)
		err, property := j.schema.GetProperty(k, className, sanitizedPropertyName)
		if err != nil {
			return q, err
		}

		janusPropertyName := string(
			j.state.getMappedPropertyName(className, sanitizedPropertyName))
		propType, err := j.schema.FindPropertyDataType(property.AtDataType)
		if err != nil {
			return q, err
		}

		if propType.IsPrimitive() {
			q, err = addPrimitivePropToQuery(q, propType, value,
				janusPropertyName, sanitizedPropertyName)
			if err != nil {
				return q, err
			}
		} else {
			result, err := j.edgesFromReferenceProp(property, value, propType, janusPropertyName, sanitizedPropertyName)
			if err != nil {
				return q, err
			}

			localEdges = append(localEdges, result.localEdges...)
			networkEdges = append(networkEdges, result.networkEdges...)
			dropTheseEdgeTypes = append(dropTheseEdgeTypes, result.edgesToDrop...)
		}
	}

	// Now drop all edges of the type we are touching
	for _, edgeLabel := range dropTheseEdgeTypes {
		q = q.Optional(gremlin.Current().OutEWithLabel(edgeLabel).HasString(PROP_REF_ID, edgeLabel).Drop())
	}

	// (Re-)Add edges to all local refs
	for _, edge := range localEdges {
		q = q.AddE(edge.PropertyName).
			FromRef("class").
			ToQuery(gremlin.G.V().HasString(PROP_UUID, edge.Reference)).
			StringProperty(PROP_REF_ID, edge.PropertyName).
			StringProperty(PROP_REF_EDGE_CREF, edge.Reference).
			StringProperty(PROP_REF_EDGE_TYPE, edge.Type).
			StringProperty(PROP_REF_EDGE_LOCATION, edge.Location)
	}

	// (Re-)Add edges to all network refs
	for _, edge := range networkEdges {
		q = q.AddE(edge.PropertyName).
			FromRef("class").
			ToQuery(
				gremlin.G.V().HasString(PROP_UUID, edge.Reference).
					Fold().
					Coalesce(gremlin.RawQuery(
						fmt.Sprintf("unfold(), addV().property(\"uuid\", \"%s\")", edge.Reference),
					)),
			).
			StringProperty(PROP_REF_ID, edge.PropertyName).
			StringProperty(PROP_REF_EDGE_CREF, edge.Reference).
			StringProperty(PROP_REF_EDGE_TYPE, edge.Type).
			StringProperty(PROP_REF_EDGE_LOCATION, edge.Location)
	}

	return q, nil
}

func addPrimitivePropToQuery(q *gremlin.Query, propType schema.PropertyDataType,
	value interface{}, janusPropertyName string, sanitizedPropertyName schema.PropertyName,
) (*gremlin.Query, error) {
	switch propType.AsPrimitive() {
	case schema.DataTypeInt:
		switch t := value.(type) {
		case int:
			q = q.Int64Property(janusPropertyName, int64(t))
		case int8:
			q = q.Int64Property(janusPropertyName, int64(t))
		case int16:
			q = q.Int64Property(janusPropertyName, int64(t))
		case int32:
			q = q.Int64Property(janusPropertyName, int64(t))
		case int64:
			q = q.Int64Property(janusPropertyName, t)
		case uint:
			q = q.Int64Property(janusPropertyName, int64(t))
		case uint8:
			q = q.Int64Property(janusPropertyName, int64(t))
		case uint16:
			q = q.Int64Property(janusPropertyName, int64(t))
		case uint32:
			q = q.Int64Property(janusPropertyName, int64(t))
		case uint64:
			q = q.Int64Property(janusPropertyName, int64(t))
		case float32:
			q = q.Int64Property(janusPropertyName, int64(t))
		case float64:
			q = q.Int64Property(janusPropertyName, int64(t))
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	case schema.DataTypeString:
		switch t := value.(type) {
		case string:
			q = q.StringProperty(janusPropertyName, t)
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	case schema.DataTypeText:
		switch t := value.(type) {
		case string:
			q = q.StringProperty(janusPropertyName, t)
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	case schema.DataTypeBoolean:
		switch t := value.(type) {
		case bool:
			q = q.BoolProperty(janusPropertyName, t)
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	case schema.DataTypeNumber:
		switch t := value.(type) {
		case float32:
			q = q.Float64Property(janusPropertyName, float64(t))
		case float64:
			q = q.Float64Property(janusPropertyName, t)
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	case schema.DataTypeDate:
		switch t := value.(type) {
		case time.Time:
			q = q.StringProperty(janusPropertyName, t.Format(time.RFC3339))
		default:
			return q, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	default:
		panic(fmt.Sprintf("Unkown primitive datatype %s", propType.AsPrimitive()))
	}

	return q, nil
}

func (j *Janusgraph) edgesFromReferenceProp(property *models.SemanticSchemaClassProperty,
	value interface{}, propType schema.PropertyDataType, janusPropertyName string, sanitizedPropertyName schema.PropertyName) (edgeFromRefProp, error) {
	result := edgeFromRefProp{}

	switch schema.CardinalityOfProperty(property) {
	case schema.CardinalityAtMostOne:
		return j.singleRef(value, propType, janusPropertyName, sanitizedPropertyName)
	case schema.CardinalityMany:
		result.edgesToDrop = []string{janusPropertyName}
		switch t := value.(type) {
		case *models.MultipleRef:
			for _, ref := range *t {
				result.localEdges = append(result.localEdges, edge{
					PropertyName: janusPropertyName,
					Reference:    ref.NrDollarCref.String(),
					Type:         ref.Type,
					Location:     *ref.LocationURL,
				})
			}
			return result, nil
		case []interface{}:
			for _, ref_ := range t {
				ref, ok := ref_.(*models.SingleRef)
				if !ok {
					return result, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
				}
				result.localEdges = append(result.localEdges, edge{
					PropertyName: janusPropertyName,
					Reference:    ref.NrDollarCref.String(),
					Type:         ref.Type,
					Location:     *ref.LocationURL,
				})
			}
			return result, nil
		default:
			return result, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
		}
	default:
		return result, fmt.Errorf("Unexpected cardinality %v",
			schema.CardinalityOfProperty(property))
	}
}

func (j *Janusgraph) singleRef(value interface{}, propType schema.PropertyDataType,
	janusPropertyName string, sanitizedPropertyName schema.PropertyName) (edgeFromRefProp, error) {
	result := edgeFromRefProp{}
	switch ref := value.(type) {
	case *models.SingleRef:
		switch ref.Type {
		case "NetworkThing", "NetworkAction":
			return j.singleNetworkRef(ref, janusPropertyName)
		case "Action", "Thing":
			return j.singleLocalRef(ref, propType, janusPropertyName, sanitizedPropertyName)
		default:
			return result, fmt.Errorf(
				"illegal value for property %s; only Thing or Action supported", ref.Type)
		}

	default:
		return result, fmt.Errorf("Illegal value for property %s", sanitizedPropertyName)
	}
}

func (j *Janusgraph) singleNetworkRef(ref *models.SingleRef, janusPropertyName string,
) (edgeFromRefProp, error) {
	result := edgeFromRefProp{}
	// We can't do any business-validation in here (such as does this
	// NetworkThing/Action really exist on that particular network instance?), as
	// we are in a (local) database connector.  Network validations are not our
	// concern. We must trust that a previous layer has verified the correctness.

	result.networkEdges = []edge{{
		PropertyName: janusPropertyName,
		Reference:    ref.NrDollarCref.String(),
		Type:         ref.Type,
		Location:     *ref.LocationURL,
	}}
	return result, nil
}

func (j *Janusgraph) singleLocalRef(ref *models.SingleRef, propType schema.PropertyDataType,
	janusPropertyName string, sanitizedPropertyName schema.PropertyName) (edgeFromRefProp, error) {
	var refClassName schema.ClassName
	result := edgeFromRefProp{}

	switch ref.Type {
	case "Action":
		var singleRefValue models.ActionGetResponse
		err := j.GetAction(nil, ref.NrDollarCref, &singleRefValue)
		if err != nil {
			return result, fmt.Errorf("Illegal value for property %s; could not resolve action with UUID: %v", ref.NrDollarCref.String(), err)
		}
		refClassName = schema.AssertValidClassName(singleRefValue.AtClass)
	case "Thing":
		var singleRefValue models.ThingGetResponse
		err := j.GetThing(nil, ref.NrDollarCref, &singleRefValue)
		if err != nil {
			return result, fmt.Errorf("Illegal value for property %s; could not resolve thing with UUID: %v", ref.NrDollarCref.String(), err)
		}
		refClassName = schema.AssertValidClassName(singleRefValue.AtClass)
	}

	// Verify the cross reference
	if !propType.ContainsClass(refClassName) {
		return result, fmt.Errorf("Illegal value for property %s; cannot point to %s", sanitizedPropertyName, ref.Type)
	}
	result.localEdges = []edge{{
		PropertyName: janusPropertyName,
		Reference:    ref.NrDollarCref.String(),
		Type:         ref.Type,
		Location:     *ref.LocationURL,
	}}

	return result, nil
}
