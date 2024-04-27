package readonly

import (
	"go/types"
	"reflect"
)

const (
	tagName            = "readonly"
	tagValueEnforceAll = "enforce_all"
)

func fieldIsEnforced(strct *types.Struct) bool {
	return allFieldsAreEnforced(strct)
}

func allFieldsAreEnforced(strct *types.Struct) bool {
	return reflect.StructTag(strct.Tag(0)).Get(tagName) == tagValueEnforceAll
}
