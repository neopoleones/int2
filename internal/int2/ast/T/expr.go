package T

type ExprVisitor interface {
	VisitBinaryExpr(*BinaryExpr) any
	VisitUnaryExpr(*UnaryExpr) any
	VisitGroupingExpr(*GroupingExpr) any
	VisitLiteralExpr(*LiteralExpr) any
	VisitVariableExpr(*VariableExpr) any
}

type Expr interface {
	Accept(ExprVisitor) any
}

type ExprStmt struct {
	NestedExpr Expr
}

func (es *ExprStmt) Accept(v StmtVisitor) any {
	return v.VisitExprStmt(es)
}

func MewExprStmt(expr Expr) *ExprStmt {
	return &ExprStmt{
		NestedExpr: expr,
	}
}
