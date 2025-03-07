package ast

import (
	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/token"
)

type Parser interface {
	ParseExpr(source []*token.Token) (T.Expr, error)
	Parse(source []*token.Token) ([]T.Stmt, error)
	Reset()
}
