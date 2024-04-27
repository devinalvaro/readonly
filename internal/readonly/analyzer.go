package readonly

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func New() *analysis.Analyzer {
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
	var filter = []ast.Node{(*ast.AssignStmt)(nil), (*ast.IncDecStmt)(nil)}

	visitor.Preorder(filter, func(node ast.Node) {
		checkAssignStmt(pass, node)
		checkIncDecStmt(pass, node)
	})

	return nil, nil
}
