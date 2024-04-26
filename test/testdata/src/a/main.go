package main

import "a/lib"

func main() {
	var strct = lib.Struct{}

	strct.A = 42 // want `readonly: field is being modified`

	strct.B = "..."
}
