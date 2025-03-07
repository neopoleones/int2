package evaluator

import (
	"reflect"

	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/interpreter"
	"github.com/destr4ct/int2/internal/int2/token"
)

type Evaluator struct {
}

func (ev *Evaluator) Evaluate(expr T.Expr) any {
	return expr.Accept(ev)
}

func (ev *Evaluator) VisitBinaryExpr(be *T.BinaryExpr) any {
	lRes := ev.Evaluate(be.Left)
	rRes := ev.Evaluate(be.Right)

	switch be.Op.Type {
	case token.TMinus:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) - rRes.(float64)
		}

	case token.TSlash:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) / rRes.(float64)
		}

	case token.TStar:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) * rRes.(float64)
		}

	case token.TPlus:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) + rRes.(float64)
		}

		if ev.suitesTypeRequirement(reflect.String, lRes, rRes) {
			return lRes.(string) + rRes.(string)
		}

	case token.TGreater:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) > rRes.(float64)
		}

	case token.TGreaterEqual:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) >= rRes.(float64)
		}

	case token.TLess:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) < rRes.(float64)
		}

	case token.TLessEqual:
		if ev.suitesTypeRequirement(reflect.Float64, lRes, rRes) {
			return lRes.(float64) <= rRes.(float64)
		}

	case token.TEqualEqual:
		return ev.areEqual(lRes, rRes)

	case token.TBangEqual:
		return !ev.areEqual(lRes, rRes)
	}

	// TODO: create error type for int2
	interpreter.RaiseBadBinary(be, lRes, rRes)
	return nil
}

func (ev *Evaluator) VisitUnaryExpr(ue *T.UnaryExpr) any {
	res := ev.Evaluate(ue.Right)

	switch ue.LeftOp.Type {
	case token.TMinus:
		if ev.suitesTypeRequirement(reflect.Float64, res) {
			return -res.(float64)
		}

	case token.TBang:
		return !ev.isTrue(res)
	}

	interpreter.RaiseBadUnary(ue, res)
	return nil
}

func (ev *Evaluator) VisitGroupingExpr(ge *T.GroupingExpr) any {
	return ev.Evaluate(ge.NestedExpr)
}

func (ev *Evaluator) VisitLiteral(le *T.LiteralExpr) any {
	return le.Value
}

func (ev *Evaluator) suitesTypeRequirement(t reflect.Kind, values ...any) bool {
	for _, v := range values {
		if reflect.ValueOf(v).Kind() != t {
			return false
		}
	}

	return true
}

func (ev *Evaluator) areEqual(v1, v2 any) bool {
	if v1 == nil && v2 == nil {
		return true
	}

	return v1 == v2
}

func (ev *Evaluator) isTrue(v any) bool {
	if v == nil {
		return false
	}

	if v, ok := v.(bool); ok {
		return v
	}

	return true
}

func Get() *Evaluator {
	return &Evaluator{}
}
