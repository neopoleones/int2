package interpreter

import (
	"fmt"
	"reflect"

	"github.com/destr4ct/int2/internal/int2/ast/T"
)

type RuntimeError struct {
	Reason string
}

func (re *RuntimeError) Error() string {
	return re.Reason
}

func Raise(reason string) { // noreturn
	panic(&RuntimeError{
		Reason: reason,
	})
}

func RaiseBadUnary(ue *T.UnaryExpr, v any) {
	op := ue.LeftOp.Lexeme

	Raise(fmt.Sprintf(
		"bad unary operation: call '%s' on '%v'", op, reflect.TypeOf(v).String(),
	))
}

func RaiseBadBinary(ue *T.BinaryExpr, v1 any, v2 any) {
	op := ue.Op.Lexeme

	Raise(fmt.Sprintf(
		"bad binary operation: call '%s' with '%v' and '%v'", op, reflect.TypeOf(v1).String(),
		reflect.TypeOf(v2).String(),
	))
}
