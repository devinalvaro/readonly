package test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/monadicEffect/readonly/internal/readonly"
)

func TestReadonly(t *testing.T) {
	var analyzer = readonly.New()

	analysistest.Run(t, analysistest.TestData(), analyzer, "a")
}
