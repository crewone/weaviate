package meta

import (
	"fmt"
	"strings"

	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/state"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
	"github.com/creativesoftwarefdn/weaviate/models"
)

type Query struct {
	params     *getmeta.Params
	nameSource nameSource
	typeSource typeSource
}

func NewQuery(params *getmeta.Params, nameSource nameSource, typeSource typeSource) *Query {
	return &Query{params: params, nameSource: nameSource, typeSource: typeSource}
}

type nameSource interface {
	GetMappedPropertyName(className schema.ClassName, propName schema.PropertyName) state.MappedPropertyName
}

type typeSource interface {
	GetProperty(kind kind.Kind, className schema.ClassName,
		propName schema.PropertyName) (error, *models.SemanticSchemaClassProperty)
	FindPropertyDataType(dataType []string) (schema.PropertyDataType, error)
}

func (b *Query) String() (string, error) {
	q := gremlin.New()

	props := b.params.Properties
	propQueries := []*gremlin.Query{}

	for _, prop := range props {
		propQuery, err := b.prop(prop)
		if err != nil {
			return "", fmt.Errorf("could not build get meta query for prop '%s': %s", prop.Name, err)
		}

		if propQuery != nil {
			propQueries = append(propQueries, propQuery)
		}
	}

	q = q.Union(propQueries...)
	return fmt.Sprintf(".%s", q.String()), nil
}

func (b *Query) prop(prop getmeta.MetaProperty) (*gremlin.Query, error) {
	if prop.Name == MetaProp {
		return b.metaProp(prop)
	}

	err, parsed := b.typeSource.GetProperty(b.params.Kind, b.params.ClassName, untitle(prop.Name))
	if err != nil {
		return nil, fmt.Errorf("could not find property '%s' in schema: %s", prop.Name, err)
	}

	dataType, err := b.typeSource.FindPropertyDataType(parsed.AtDataType)
	if err != nil {
		return nil, fmt.Errorf("could not find data type of prop '%s': %s", prop.Name, err)
	}

	if !dataType.IsPrimitive() {
		// skip, as we can get all info for ref-props from the TypeInspector
		return nil, nil
	}

	switch dataType.AsPrimitive() {
	case schema.DataTypeBoolean:
		return b.booleanProp(prop)
	case schema.DataTypeString, schema.DataTypeDate:
		return b.stringProp(prop)
	case schema.DataTypeInt, schema.DataTypeNumber:
		return b.intProp(prop)
	default:
		return nil, fmt.Errorf("unsupported primitive property data type: %s", dataType.AsPrimitive())
	}
}

func (b *Query) mappedPropertyName(className schema.ClassName,
	propName schema.PropertyName) string {
	if b.nameSource == nil {
		return string(propName)
	}

	return string(b.nameSource.GetMappedPropertyName(className, propName))
}

func untitle(propName schema.PropertyName) schema.PropertyName {
	asString := string(propName)
	return schema.PropertyName(strings.ToLower(string(asString[0])) + asString[1:])
}
