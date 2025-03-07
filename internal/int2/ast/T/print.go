package T

type PrintStmt struct {
	NestedExpr Expr
}

func (ps *PrintStmt) Accept(v StmtVisitor) any {
	return v.VisitPrintStmt(ps)
}

func MewPrintStmt(expr Expr) *PrintStmt {
	return &PrintStmt{
		NestedExpr: expr,
	}
}
