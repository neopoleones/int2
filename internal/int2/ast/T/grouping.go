package T

type GroupingExpr struct {
	NestedExpr Expr
}

func (gr *GroupingExpr) Accept(v ExprVisitor) any {
	return v.VisitGroupingExpr(gr)
}

func NewGroupingExpr(nested Expr) *GroupingExpr {
	return &GroupingExpr{
		NestedExpr: nested,
	}
}
