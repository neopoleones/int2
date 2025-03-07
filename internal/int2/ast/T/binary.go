package T

import (
	"github.com/destr4ct/int2/internal/int2/token"
)

type BinaryExpr struct {
	Left  Expr
	Right Expr
	Op    *token.Token
}

func (b *BinaryExpr) Accept(v ExprVisitor) any {
	return v.VisitBinaryExpr(b)
}

func NewBinaryExpr(left Expr, right Expr, op *token.Token) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Right: right,
		Op:    op,
	}
}
