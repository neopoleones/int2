package scanner

import (
	"errors"
	"fmt"

	"github.com/destr4ct/int2/internal/int2/token"
)

var (
	ErrLexStringUnterminated = errors.New("unterminated string specified")
	ErrLexBadNumber          = errors.New("incorrect Number specified")
	ErrLexIncorrectChar      = errors.New("incorrect character being used in source code")
)

type lexError struct {
	err error
	tok *token.Token
}

func (le *lexError) Error() string {
	return fmt.Sprintf("line %d at '%v': %s", le.tok.Line, le.tok.Lexeme, le.err)
}

func withContext(err error, tok *token.Token) *lexError {
	return &lexError{
		err: err,
		tok: tok,
	}
}

func Validate(tokens []*token.Token) error {
	for _, tok := range tokens {
		if tok.LexErr != nil {
			return withContext(tok.LexErr, tok)
		}
	}

	return nil
}
