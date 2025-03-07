package T

type StmtVisitor interface {
	VisitPrintStmt(*PrintStmt) any
	VisitExprStmt(*ExprStmt) any
}

type Stmt interface {
	Accept(StmtVisitor) any
}
