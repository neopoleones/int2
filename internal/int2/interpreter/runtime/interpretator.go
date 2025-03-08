package runtime

import (
	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/interpreter"
	"github.com/destr4ct/int2/internal/int2/interpreter/evaluator"
)

type Interpretator struct {
	evaluator.Evaluator
}

func (int2 *Interpretator) Execute(stmt T.Stmt) {
	_ = stmt.Accept(int2)
}

func (int2 *Interpretator) Interpret(statements []T.Stmt) {
	for _, stmt := range statements {
		int2.Execute(stmt)
	}
}

func Get(storage interpreter.Environ) *Interpretator {
	return &Interpretator{
		Evaluator: *evaluator.Get(storage),
	}
}
