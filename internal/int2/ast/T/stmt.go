package T

type StmtVisitor interface {
	VisitPrintStmt(*PrintStmt) any
	VisitExprStmt(*ExprStmt) any
	VisitVarStmt(*VarStmt) any
}

type Stmt interface {
	Accept(StmtVisitor) any
}
