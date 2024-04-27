package readonly

import (
	"go/types"
	"reflect"
)

func findStructField(strct *types.Struct, fieldName string) (field, bool) {
	for i := 0; i < strct.NumFields(); i++ {
		var structField = strct.Field(i)

		if structField.Name() == fieldName {
			return field{
				Name: fieldName,
				Tag:  strct.Tag(i),
			}, true
		}
	}

	return field{}, false
}

const (
	tagName         = "readonly"
	tagValueIgnore  = "ignore"
)

type field struct {
	Name string
	Tag  string
}

func (s field) isIgnored() bool {
	return reflect.StructTag(s.Tag).Get(tagName) == tagValueIgnore
}
