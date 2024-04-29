package readonly

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

type Checker struct {
	pass *analysis.Pass
}

func NewChecker(pass *analysis.Pass) Checker {
	return Checker{
		pass: pass,
	}
}

func (s Checker) Nodes() []ast.Node {
	return []ast.Node{(*ast.AssignStmt)(nil), (*ast.IncDecStmt)(nil)}
}

func (s Checker) Check(node ast.Node) {
	s.checkAssignStmt(node)
	s.checkIncDecStmt(node)
}

func (s Checker) checkAssignStmt(node ast.Node) {
	var assignStmt, ok = node.(*ast.AssignStmt)
	if !ok {
		return
	}

	for _, expr := range assignStmt.Lhs {
		s.checkLhs(expr)
	}
}

func (s Checker) checkIncDecStmt(node ast.Node) {
	var incDecStmt, ok = node.(*ast.IncDecStmt)
	if !ok {
		return
	}

	s.checkLhs(incDecStmt.X)
}

func (s Checker) checkLhs(expr ast.Expr) {
	s.checkSelector(expr)
	s.checkStar(expr)
	s.checkIndex(expr)
	s.checkParen(expr)
}

func (s Checker) checkSelector(expr ast.Expr) {
	var selector, ok = expr.(*ast.SelectorExpr)
	if !ok {
		return
	}

	var fieldSelector = selector.Sel
	if s.pass.TypesInfo.ObjectOf(fieldSelector).Pkg() == s.pass.Pkg {
		return
	}
	// invariant: The field is selected outside of its package.

	var selectedType = s.pass.TypesInfo.TypeOf(selector.X)
	var selectedStruct *types.Struct
	if selectedStruct, ok = selectedType.Underlying().(*types.Struct); !ok {
		return
	}

	if fieldIsEnforced(selectedStruct) {
		s.pass.Reportf(fieldSelector.Pos(), "readonly: field is being modified")
	}
}

func (s Checker) checkStar(expr ast.Expr) {
	if starExpr, ok := expr.(*ast.StarExpr); ok {
		s.checkLhs(starExpr.X)
	}
}

func (s Checker) checkIndex(expr ast.Expr) {
	if indexExpr, ok := expr.(*ast.IndexExpr); ok {
		s.checkLhs(indexExpr.X)
	}
}

func (s Checker) checkParen(expr ast.Expr) {
	var paren, ok = expr.(*ast.ParenExpr)
	if !ok {
		return
	}

	s.checkLhs(paren.X)
}
