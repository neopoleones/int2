package T

import "github.com/destr4ct/int2/internal/int2/token"

type VarStmt struct {
	Name        *token.Token
	Initializer Expr
}

func (vs *VarStmt) Accept(v StmtVisitor) any {
	return v.VisitVarStmt(vs)
}

type VariableExpr struct {
	Name *token.Token
}

func (ve *VariableExpr) Accept(v ExprVisitor) any {
	return v.VisitVariableExpr(ve)
}

func NewVarStmt(name *token.Token, initializer Expr) *VarStmt {
	return &VarStmt{
		Name:        name,
		Initializer: initializer,
	}
}

func NewVariableExpr(name *token.Token) *VariableExpr {
	return &VariableExpr{
		Name: name,
	}
}
