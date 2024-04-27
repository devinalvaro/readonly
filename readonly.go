package readonly

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/devinalvaro/readonly/internal/readonly"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "readonly",
		Doc:      "make exported fields read-only from outside the package",
		URL:      "https://github.com/devinalvaro/readonly",
		Flags:    flag.FlagSet{},
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	var visitor = pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	var checker = readonly.NewChecker(pass)

	visitor.Preorder(checker.Nodes(), checker.Check)

	return nil, nil
}
