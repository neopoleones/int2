package rdparser

import (
	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/token"
)

type RecursiveDescentParser struct {
	curr uint64

	sourceTokens []*token.Token
}

func (rdp *RecursiveDescentParser) ParseExpr(source []*token.Token) (T.Expr, error) {
	defer rdp.Reset()
	rdp.sourceTokens = source

	// Unrecoverable error goes linked
	return rdp.parseExpression()
}

func (rdp *RecursiveDescentParser) Parse(source []*token.Token) ([]T.Stmt, error) {
	statements := make([]T.Stmt, 0)

	for rdp.hasNextToken() {
		stmt, err := rdp.parseStatement()
		if err != nil {
			return statements, err
		}

		statements = append(statements, stmt)
	}

	return statements, nil
}

func (rdp *RecursiveDescentParser) Reset() {
	rdp.curr = 0
}

func (rdp *RecursiveDescentParser) parseStatement() (T.Stmt, error) {
	if rdp.headMatchType(token.TPrint) {
		return rdp.parsePrintStatement()
	}

	return rdp.parseExprStatement()
}

func (rdp *RecursiveDescentParser) parsePrintStatement() (T.Stmt, error) {
	return nil, nil
}

func (rdp *RecursiveDescentParser) parseExprStatement() (T.Stmt, error) {
	return nil, nil
}

func (rdp *RecursiveDescentParser) parseExpression() (T.Expr, error) {
	return rdp.parseEqualityExpr()
}

func (rdp *RecursiveDescentParser) parseEqualityExpr() (T.Expr, error) {
	expr, err := rdp.parseComparisonExpr()
	if err != nil {
		return nil, err
	}

	for rdp.headMatchType(token.TBangEqual, token.TEqualEqual) {
		op := rdp.previousToken()
		right, err := rdp.parseComparisonExpr()
		if err != nil {
			return nil, err
		}

		expr = T.NewBinaryExpr(expr, right, op)
	}

	return expr, nil
}

func (rdp *RecursiveDescentParser) parseComparisonExpr() (T.Expr, error) {
	expr, err := rdp.parseTermExpr()
	if err != nil {
		return nil, err
	}

	for rdp.headMatchType(token.TGreater, token.TLess, token.TGreaterEqual, token.TLessEqual) {
		op := rdp.previousToken()
		right, err := rdp.parseTermExpr()
		if err != nil {
			return nil, err
		}

		expr = T.NewBinaryExpr(expr, right, op)
	}

	return expr, nil
}

func (rdp *RecursiveDescentParser) parseTermExpr() (T.Expr, error) {
	expr, err := rdp.parseFactorExpr()
	if err != nil {
		return nil, err
	}

	for rdp.headMatchType(token.TMinus, token.TPlus) {
		op := rdp.previousToken()
		right, err := rdp.parseFactorExpr()
		if err != nil {
			return nil, err
		}

		expr = T.NewBinaryExpr(expr, right, op)
	}

	return expr, nil
}

func (rdp *RecursiveDescentParser) parseFactorExpr() (T.Expr, error) {
	expr, err := rdp.parseUnaryExpr()
	if err != nil {
		return nil, err
	}

	for rdp.headMatchType(token.TSlash, token.TStar) {
		op := rdp.previousToken()
		right, err := rdp.parseUnaryExpr()
		if err != nil {
			return nil, err
		}

		expr = T.NewBinaryExpr(expr, right, op)
	}

	return expr, nil
}

func (rdp *RecursiveDescentParser) parseUnaryExpr() (T.Expr, error) {
	if rdp.headMatchType(token.TMinus, token.TBang) {
		op := rdp.previousToken()
		expr, err := rdp.parseUnaryExpr()
		if err != nil {
			return nil, err
		}

		return T.NewUnaryExpr(op, expr), nil
	}

	return rdp.parsePrimaryExpr()
}

func (rdp *RecursiveDescentParser) parsePrimaryExpr() (T.Expr, error) {
	if rdp.headMatchType(token.TFalse) {
		return T.NewLiteralExpr(false), nil
	}

	if rdp.headMatchType(token.TTrue) {
		return T.NewLiteralExpr(true), nil
	}

	if rdp.headMatchType(token.TNil) {
		return T.NewLiteralExpr(nil), nil
	}

	if rdp.headMatchType(token.TNumber, token.TString) {
		return T.NewLiteralExpr(rdp.previousToken().Literal), nil
	}

	if rdp.headMatchType(token.TLParent) {
		expr, err := rdp.parseExpression()
		if err != nil {
			return nil, err
		}

		if _, err := rdp.consumeConcrete(token.TRParent, "unclosed parent"); err != nil {
			return nil, err
		}

		return T.NewGroupingExpr(expr), nil
	}

	return nil, newASTError(rdp.headToken(), "expected expression")
}

func (rdp *RecursiveDescentParser) headMatchType(tokTypes ...token.TokenType) bool {
	for _, tokType := range tokTypes {
		if rdp.checkTokType(tokType) {
			_ = rdp.consumeToken()
			return true
		}
	}

	return false
}

func (rdp *RecursiveDescentParser) checkTokType(tokType token.TokenType) bool {
	if !rdp.hasNextToken() {
		return false
	}

	return rdp.headToken().Type == tokType
}

func (rdp *RecursiveDescentParser) consumeConcrete(tokType token.TokenType, errMsg string) (*token.Token, error) {
	tok := rdp.consumeToken()

	if tokType != tok.Type {
		return nil, newASTError(tok, errMsg)
	}

	return tok, nil
}

func (rdp *RecursiveDescentParser) consumeToken() *token.Token {
	if rdp.hasNextToken() {
		rdp.curr++
	}

	return rdp.previousToken()
}

func (rdp *RecursiveDescentParser) synchronize() {
	_ = rdp.consumeToken()
	for rdp.hasNextToken() {
		if rdp.previousToken().Type == token.TSemicolon {
			return
		}

		switch rdp.headToken().Type {
		case token.TClass, token.TFun,
			token.TVar, token.TFor, token.TIf, token.TWhile,
			token.TPrint, token.TReturn:

			return
		}

		_ = rdp.consumeToken()
	}
}

func (rdp *RecursiveDescentParser) hasNextToken() bool {
	return rdp.headToken().Type != token.TEOF
}

func (rdp *RecursiveDescentParser) headToken() *token.Token {
	return rdp.sourceTokens[rdp.curr]
}

func (rdp *RecursiveDescentParser) previousToken() *token.Token {
	return rdp.sourceTokens[rdp.curr-1]
}

func Get() *RecursiveDescentParser {
	return &RecursiveDescentParser{
		curr: 0,
	}
}
