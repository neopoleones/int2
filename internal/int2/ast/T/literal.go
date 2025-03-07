package T

type LiteralExpr struct {
	Value any
}

func (li *LiteralExpr) Accept(v ExprVisitor) any {
	return v.VisitLiteral(li)
}

func NewLiteralExpr(v any) *LiteralExpr {
	return &LiteralExpr{
		Value: v,
	}
}
