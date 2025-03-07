package token

import "fmt"

type Token struct {
	Type TokenType

	Lexeme  string
	Literal any

	Line uint64

	LexErr error
}

func PrintTokens(tokens []*Token) {
	fmt.Println("Tokens:")
	for _, tok := range tokens {
		fmt.Println("\t", tok)
	}
}

func (tok *Token) String() string {
	// If keyword
	if v, found := revertedTranslation[tok.Type]; found {
		return fmt.Sprintf("Keyword [line: %d, k: %s, e: %+v]", tok.Line, v, tok.LexErr)
	} else if tok.Type == TIdentifier {
		return fmt.Sprintf("Identifier [line: %d, i: %s, e: %+v]", tok.Line, tok.Lexeme, tok.LexErr)
	} else if tok.Type == TString {
		return fmt.Sprintf("String [line: %d, l: %v, e: %+v]", tok.Line, tok.Literal, tok.LexErr)
	} else if tok.Type == TNumber {
		return fmt.Sprintf("Number [line: %d, l: %v, e: %+v]", tok.Line, tok.Literal, tok.LexErr)
	} else if tok.Type == TEOF {
		return fmt.Sprintf("EOF [line: %d]", tok.Line)
	} else if tok.Type == TBadChar {
		return fmt.Sprintf("BADCHAR [line: %d, l: %s, e: %+v]", tok.Line, tok.Lexeme, tok.LexErr)
	} else {
		return fmt.Sprintf("Operator [line: %d, o: %s, e: %+v]", tok.Line, tok.Lexeme, tok.LexErr)
	}
}

func (tok *Token) WithError(err error) *Token {
	tok.LexErr = err
	return tok
}

func New(tokType TokenType, lexeme string, literal any, line uint64) *Token {
	return &Token{
		Type:    tokType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
