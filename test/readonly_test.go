package test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/devinalvaro/readonly"
)

func TestReadonly(t *testing.T) {
	var analyzer = readonly.NewAnalyzer()
	analysistest.Run(t, analysistest.TestData(), analyzer, "noflags")
}
