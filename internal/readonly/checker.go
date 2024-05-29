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
	switch node := node.(type) {
	case *ast.AssignStmt:
		s.checkAssignStmt(node)
	case *ast.IncDecStmt:
		s.checkIncDecStmt(node)
	}
}

func (s Checker) checkAssignStmt(assign *ast.AssignStmt) {
	for _, expr := range assign.Lhs {
		s.checkLhs(expr)
	}
}

func (s Checker) checkIncDecStmt(incDec *ast.IncDecStmt) {
	s.checkLhs(incDec.X)
}

func (s Checker) checkLhs(lhs ast.Expr) {
	switch expr := lhs.(type) {
	case *ast.SelectorExpr:
		s.checkSelectorLhs(expr)
	case *ast.StarExpr:
		s.checkStarLhs(expr)
	case *ast.IndexExpr:
		s.checkIndexLhs(expr)
	case *ast.ParenExpr:
		s.checkParenLhs(expr)
	}
}

func (s Checker) checkSelectorLhs(selector *ast.SelectorExpr) {
	var fieldSelector = selector.Sel
	if s.pass.TypesInfo.ObjectOf(fieldSelector).Pkg() == s.pass.Pkg {
		return
	}
	// invariant: The field is selected outside of its package.

	var selectedType = s.pass.TypesInfo.TypeOf(selector.X)
	var selectedStruct, ok = assertStructType(selectedType)
	if !ok {
		return
	}

	if fieldIsEnforced(selectedStruct, fieldSelector.Name) {
		s.pass.Reportf(fieldSelector.Pos(), "readonly: field is being modified")
	}
}

func (s Checker) checkStarLhs(star *ast.StarExpr) {
	s.checkLhs(star.X)
}

func (s Checker) checkIndexLhs(index *ast.IndexExpr) {
	s.checkLhs(index.X)
}

func (s Checker) checkParenLhs(paren *ast.ParenExpr) {
	s.checkLhs(paren.X)
}

func assertStructType(typ types.Type) (*types.Struct, bool) {
	if pointer, ok := typ.Underlying().(*types.Pointer); ok {
		return assertStructType(pointer.Elem())
	}

	var strct, ok = typ.Underlying().(*types.Struct)
	return strct, ok
}
