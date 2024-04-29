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

	var newValue = true
	strct.Pointer = &newValue // want `readonly: field is being modified`
	*strct.Pointer = newValue // want `readonly: field is being modified`

	strct.NestedPointer = &strct.Pointer   // want `readonly: field is being modified`
	*strct.NestedPointer = strct.Pointer   // want `readonly: field is being modified`
	**strct.NestedPointer = *strct.Pointer // want `readonly: field is being modified`
}
