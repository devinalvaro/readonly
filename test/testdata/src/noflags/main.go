package main

import "noflags/lib"

func main() {
	var strct = lib.Struct{}

	strct.Number = 42  // want `readonly: field is being modified`
	strct.Number += 1  // want `readonly: field is being modified`
	strct.Number -= 1  // want `readonly: field is being modified`
	strct.Number *= 1  // want `readonly: field is being modified`
	strct.Number /= 1  // want `readonly: field is being modified`
	strct.Number %= 1  // want `readonly: field is being modified`
	strct.Number &= 1  // want `readonly: field is being modified`
	strct.Number |= 1  // want `readonly: field is being modified`
	strct.Number ^= 1  // want `readonly: field is being modified`
	strct.Number <<= 1 // want `readonly: field is being modified`
	strct.Number >>= 1 // want `readonly: field is being modified`
	strct.Number &^= 1 // want `readonly: field is being modified`
	strct.Number++     // want `readonly: field is being modified`
	strct.Number--     // want `readonly: field is being modified`

	strct.Pointer = nil   // want `readonly: field is being modified`
	*strct.Pointer = true // want `readonly: field is being modified`

	strct.NestedPointer = &strct.Pointer   // want `readonly: field is being modified`
	*strct.NestedPointer = strct.Pointer   // want `readonly: field is being modified`
	**strct.NestedPointer = *strct.Pointer // want `readonly: field is being modified`

	strct.Slice = make([]string, 0) // want `readonly: field is being modified`
	strct.Slice[0] = "any"          // want `readonly: field is being modified`

	strct.SlicePointer = &strct.Slice // want `readonly: field is being modified`
	(*strct.SlicePointer)[0] = "any"  // want `readonly: field is being modified`

	strct.NestedSlice = make([][]string, 0)  // want `readonly: field is being modified`
	strct.NestedSlice[0] = make([]string, 0) // want `readonly: field is being modified`
	strct.NestedSlice[0][1] = "any"          // want `readonly: field is being modified`

	strct.Map = make(map[string]struct{}) // want `readonly: field is being modified`
	strct.Map["any"] = struct{}{}         // want `readonly: field is being modified`

	strct.MapPointer = &strct.Map           // want `readonly: field is being modified`
	(*strct.MapPointer)["any"] = struct{}{} // want `readonly: field is being modified`

	strct.NestedMap = make(map[string]map[string]struct{}) // want `readonly: field is being modified`
	strct.NestedMap["any"] = make(map[string]struct{})     // want `readonly: field is being modified`
	strct.NestedMap["any"]["thing"] = struct{}{}           // want `readonly: field is being modified`

	strct.Self.Number = 42 // want `readonly: field is being modified`

	var outer = lib.Outer{}

	outer.Struct.Number = 42 // want `readonly: field is being modified`
}
