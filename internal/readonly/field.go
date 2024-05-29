package readonly

import (
	"go/types"
	"reflect"
)

const (
	tagName            = "readonly"
	tagValueEnforce    = "enforce"
	tagValueEnforceAll = "enforce_all"
)

func fieldIsEnforced(strct *types.Struct, fieldName string) bool {
	if allFieldsAreEnforced(strct) {
		return true
	}

	var fieldTag, ok = findStructFieldTag(strct, fieldName)
	if !ok {
		return false
	}

	return reflect.StructTag(fieldTag).Get(tagName) == tagValueEnforce
}

func allFieldsAreEnforced(strct *types.Struct) bool {
	return reflect.StructTag(strct.Tag(0)).Get(tagName) == tagValueEnforceAll
}

func findStructFieldTag(strct *types.Struct, fieldName string) (string, bool) {
	for i := 0; i < strct.NumFields(); i++ {
		var structField = strct.Field(i)

		if structField.Name() == fieldName {
			return strct.Tag(i), true
		}
	}

	return "", false
}
