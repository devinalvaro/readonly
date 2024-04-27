package main

import "a/lib"

func main() {
	var strct = lib.Struct{}

	strct.A = 42 // want `readonly: field is being modified`
	strct.A += 1 // want `readonly: field is being modified`
	strct.A *= 1 // want `readonly: field is being modified`
	strct.A++    // want `readonly: field is being modified`
	strct.A--    // want `readonly: field is being modified`

	strct.B = 42
	strct.B += 1
	strct.B *= 1
	strct.B++
	strct.B--
}
