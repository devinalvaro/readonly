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

func checkSelector(pass *analysis.Pass, expr ast.Expr) {
	var selector, ok = expr.(*ast.SelectorExpr)
	if !ok {
		return
	}

	var selectorField = pass.TypesInfo.ObjectOf(selector.Sel)
	if selectorField.Pkg() == pass.Pkg {
		return
	}
	// invariant: The field is selected outside of its package.

	var selectorType = pass.TypesInfo.TypeOf(selector.X)
	var selectorStruct *types.Struct
	if selectorStruct, ok = selectorType.Underlying().(*types.Struct); !ok {
		return
	}

	if field, ok := findStructField(selectorStruct, selectorField.Name()); ok {
		if !field.isIgnored() {
			pass.Reportf(selectorField.Pos(), "readonly: field is being modified")
		}
	}
}
