package schema

import (
	"sort"
	"strings"

	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	libschema "github.com/semi-technologies/weaviate/entities/schema"
)

// RefFinder is a helper that lists classes and their possible paths to to a
// desired target class.
//
// For example if the target class is "car". It might list:
// - Person, drives, Car
// - Person, owns, Car
// - Person, friendsWith, Person, drives, Car
// etc.
//
// It will stop at a preconfigured depth limit, to avoid infinite results, such
// as:
// - Person, friendsWith, Person, friendsWith, Person, ..., drives Car
type RefFinder struct {
	schemaGetter schemaGetterForRefFinder
	depthLimit   int
}

// NewRefFinder with SchemaGetter and depth limit
func NewRefFinder(getter schemaGetterForRefFinder, depthLimit int) *RefFinder {
	return &RefFinder{
		schemaGetter: getter,
		depthLimit:   depthLimit,
	}
}

type schemaGetterForRefFinder interface {
	GetSchemaSkipAuth() libschema.Schema
}

func (r *RefFinder) Find(className libschema.ClassName) []filters.Path {
	schema := r.schemaGetter.GetSchemaSkipAuth()

	var classes []*models.Class
	if schema.Actions != nil {
		classes = append(classes, schema.Actions.Classes...)
	}
	if schema.Things != nil {
		classes = append(classes, schema.Things.Classes...)
	}

	return r.findInClassList(className, classes, schema)
}

func (r *RefFinder) findInClassList(needle libschema.ClassName, classes []*models.Class,
	schema libschema.Schema) []filters.Path {
	var out []filters.Path

	for _, class := range classes {
		path, ok := r.hasRefTo(needle, class, schema)
		if !ok {
			continue
		}

		out = append(out, path...)
	}

	return r.sortByPathLen(out)
}

func (r *RefFinder) hasRefTo(needle libschema.ClassName, class *models.Class,
	schema libschema.Schema) ([]filters.Path, bool) {
	var out []filters.Path

	for _, prop := range class.Properties {
		dt, err := schema.FindPropertyDataType(prop.DataType)
		if err != nil {
			// silently ignore, maybe the property was deleted in the meantime
		}

		if dt.IsPrimitive() {
			continue
		}

		for _, haystack := range dt.Classes() {
			refs := r.refsPerClass(needle, class, prop.Name, haystack, schema)
			out = append(out, refs...)

		}
	}

	return out, len(out) > 0
}

func (r *RefFinder) refsPerClass(needle libschema.ClassName, class *models.Class,
	propName string, haystack libschema.ClassName, schema libschema.Schema) []filters.Path {
	if haystack == needle {
		// direct match
		return []filters.Path{
			filters.Path{
				Class:    libschema.ClassName(class.Class),
				Property: libschema.PropertyName(strings.Title(propName)),
				Child: &filters.Path{
					Class:    needle,
					Property: "uuid",
				},
			},
		}
	}

	// could still be an indirect (recursive) match
	innerClass := schema.FindClassByName(haystack)
	if innerClass == nil {
		return nil
	}
	paths, ok := r.hasRefTo(needle, innerClass, schema)
	if !ok {
		return nil
	}

	var out []filters.Path
	for _, path := range paths {
		out = append(out, filters.Path{
			Class:    libschema.ClassName(class.Class),
			Property: libschema.PropertyName(strings.Title(propName)),
			Child:    &path,
		})
	}

	return out
}

func (r *RefFinder) sortByPathLen(in []filters.Path) []filters.Path {
	sort.Slice(in, func(i, j int) bool {
		return len(in[i].Slice()) < len(in[j].Slice())
	})

	return in
}
