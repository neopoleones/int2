package rdparser

import (
	"fmt"

	"github.com/destr4ct/int2/internal/int2/token"
)

type ASTError struct {
	causee *token.Token
	helper string

	backErr *ASTError
}

func (err *ASTError) Error() string {
	baseErr := fmt.Sprintf("Line %d, at '%v': %s", err.causee.Line, err.causee.Lexeme, err.helper)

	if err.backErr != nil {
		baseErr = fmt.Sprintf("%s\n%s", baseErr, err.backErr.Error())
	}

	return baseErr
}

func newASTError(causee *token.Token, helper string, we ...*ASTError) *ASTError {
	err := &ASTError{
		causee: causee,
		helper: helper,
	}

	if len(we) == 1 {
		err.backErr = we[0]
	}

	return err
}
