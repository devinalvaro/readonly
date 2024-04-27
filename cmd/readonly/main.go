package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/devinalvaro/readonly"
)

func main() {
	singlechecker.Main(readonly.NewAnalyzer())
}
