package lib

type Struct struct {
	Number        int `readonly:"enforce_all"`
	Pointer       *bool
	NestedPointer **bool
	Slice         []string
	Map           map[string]struct{}
	SlicePointer  *[]string
	MapPointer    *map[string]struct{}
}
