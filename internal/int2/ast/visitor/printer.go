package visitor

import (
	"strconv"
	"strings"

	"github.com/destr4ct/int2/internal/int2/ast/T"
)

/*
	VisitBinaryExpr(*BinaryExpr) any
	VisitUnaryExpr(*UnaryExpr) any
	VisitGroupingExpr(*GroupingExpr) any
	VisitLiteral(*LiteralExpr) any
*/

type AstPrinter struct {
}

func (ap *AstPrinter) wrap(name string, expressionList ...T.Expr) string {
	var sb strings.Builder

	sb.WriteByte('(')
	sb.WriteString(name)

	for _, expr := range expressionList {
		sb.WriteByte(' ')
		sb.WriteString(expr.Accept(ap).(string))
	}

	sb.WriteByte(')')
	return sb.String()
}

func (ap *AstPrinter) Stringify(stmt T.Stmt) string {
	return stmt.Accept(ap).(string)
}

func (ap *AstPrinter) VisitBinaryExpr(be *T.BinaryExpr) any {
	return ap.wrap(be.Op.Lexeme, be.Left, be.Right)
}

func (ap *AstPrinter) VisitUnaryExpr(be *T.UnaryExpr) any {
	return ap.wrap(be.LeftOp.Lexeme, be.Right)
}

func (ap *AstPrinter) VisitGroupingExpr(be *T.GroupingExpr) any {
	return ap.wrap("group", be.NestedExpr)
}

func (ap *AstPrinter) VisitLiteral(be *T.LiteralExpr) any {
	// Literal is value of [nil, float64, string]
	if num, ok := be.Value.(float64); ok {
		return strconv.FormatFloat(num, 'f', -1, 64)
	}

	if v, ok := be.Value.(bool); ok {
		return strconv.FormatBool(v)
	}

	if be.Value == nil {
		return "nil"
	}

	return be.Value
}

func (ap *AstPrinter) VisitPrintStmt(stmt *T.PrintStmt) any {
	return ap.wrap("print", stmt.NestedExpr)
}

func (ap *AstPrinter) VisitExprStmt(stmt *T.ExprStmt) any {
	return ap.wrap("expr", stmt.NestedExpr)
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}
