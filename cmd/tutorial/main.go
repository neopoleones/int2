package main

import (
	"fmt"

	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/ast/visitor"
	"github.com/destr4ct/int2/internal/int2/token"
)

func main() {
	expr := T.NewBinaryExpr(
		T.NewUnaryExpr(
			token.New(token.TMinus, "-", nil, 0),
			T.NewLiteralExpr(127.0),
		),
		T.NewGroupingExpr(
			T.NewLiteralExpr(20.1),
		),
		token.New(token.TStar, "*", nil, 0),
	)

	ap := visitor.NewAstPrinter()
	fmt.Println(ap.Stringify(expr))
}
