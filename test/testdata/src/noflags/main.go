package main

import "noflags/lib"

func main() {
	var strct = lib.Struct{}

	strct.A = 42 // want `readonly: field is being modified`
	strct.A += 1 // want `readonly: field is being modified`
	strct.A *= 1 // want `readonly: field is being modified`
	strct.A++    // want `readonly: field is being modified`
	strct.A--    // want `readonly: field is being modified`

	strct.B = 42 // want `readonly: field is being modified`
	strct.B += 1 // want `readonly: field is being modified`
	strct.B *= 1 // want `readonly: field is being modified`
	strct.B++    // want `readonly: field is being modified`
	strct.B--    // want `readonly: field is being modified`
}
