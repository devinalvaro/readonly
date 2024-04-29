package lib

type Struct struct {
	Number int `readonly:"enforce_all"`

	Pointer       *bool
	NestedPointer **bool

	Slice        []string
	SlicePointer *[]string
	NestedSlice  [][]string

	Map        map[string]struct{}
	MapPointer *map[string]struct{}
	NestedMap  map[string]map[string]struct{}

	Self *Struct
}

type Outer struct {
	Struct Struct
}
