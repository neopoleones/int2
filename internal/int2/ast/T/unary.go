package T

import "github.com/destr4ct/int2/internal/int2/token"

type UnaryExpr struct {
	LeftOp *token.Token
	Right  Expr
}

func (u *UnaryExpr) Accept(v ExprVisitor) any {
	return v.VisitUnaryExpr(u)
}

func NewUnaryExpr(left *token.Token, right Expr) *UnaryExpr {
	return &UnaryExpr{
		LeftOp: left,
		Right:  right,
	}
}
