package readonly

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func checkAssignStmt(pass *analysis.Pass, node ast.Node) {
	var assignStmt, ok = node.(*ast.AssignStmt)
	if !ok {
		return
	}

	for _, expr := range assignStmt.Lhs {
		checkSelector(pass, expr)
	}
}

func checkIncDecStmt(pass *analysis.Pass, node ast.Node) {
	var incDecStmt, ok = node.(*ast.IncDecStmt)
	if !ok {
		return
	}

	checkSelector(pass, incDecStmt.X)
}

func checkSelector(pass *analysis.Pass, expr ast.Expr) {
	var selector, ok = expr.(*ast.SelectorExpr)
	if !ok {
		return
	}

	var fieldSelector = selector.Sel
	if pass.TypesInfo.ObjectOf(fieldSelector).Pkg() == pass.Pkg {
		return
	}
	// invariant: The field is selected outside of its package.

	var selectedType = pass.TypesInfo.TypeOf(selector.X)
	var selectedStruct *types.Struct
	if selectedStruct, ok = selectedType.Underlying().(*types.Struct); !ok {
		return
	}

	if field, ok := findStructField(selectedStruct, fieldSelector.Name); ok {
		if !field.isIgnored() {
			pass.Reportf(fieldSelector.Pos(), "readonly: field is being modified")
		}
	}
}
