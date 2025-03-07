package interpreter

import "github.com/destr4ct/int2/internal/int2/ast/T"

type Interpreter interface {
	T.ExprVisitor

	Evaluate(expr T.Expr) any
}
