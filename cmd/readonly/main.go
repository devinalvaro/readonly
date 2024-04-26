package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/devinalvaro/readonly/internal/readonly"
)

func main() {
	singlechecker.Main(readonly.New())
}
