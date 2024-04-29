package lib

type Struct struct {
	Number        int `readonly:"enforce_all"`
	Pointer       *bool
	NestedPointer **bool
}
