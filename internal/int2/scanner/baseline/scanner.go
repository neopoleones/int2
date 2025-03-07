package baseline

import (
	"math"
	"strconv"

	"github.com/destr4ct/int2/internal/int2/scanner"
	"github.com/destr4ct/int2/internal/int2/token"
)

type BaselineScanner struct {
	start   uint64
	current uint64

	currLine uint64

	rawSources string
}

func (sc *BaselineScanner) hasNextToken() bool {
	return sc.current < uint64(len(sc.rawSources))
}

func (sc *BaselineScanner) check(b byte) bool {
	if sc.hasNextToken() && sc.rawSources[sc.current] == b {
		sc.current++
		return true
	}

	return false
}

func (sc *BaselineScanner) lookahead() byte {
	if sc.hasNextToken() {
		return sc.rawSources[sc.current]
	}

	return '\x00'
}

func (sc *BaselineScanner) scanNext() *token.Token {
	var tok *token.Token

	// Scan for easy tokens
	b := sc.consumeByte()
	switch b {
	case '(':
		tok = sc.instantiateToken(token.TLParent, nil)
	case ')':
		tok = sc.instantiateToken(token.TRParent, nil)
	case '{':
		tok = sc.instantiateToken(token.TLBrace, nil)
	case '}':
		tok = sc.instantiateToken(token.TRBrace, nil)
	case '.':
		tok = sc.instantiateToken(token.TDot, nil)
	case ',':
		tok = sc.instantiateToken(token.TComma, nil)
	case '-':
		tok = sc.instantiateToken(token.TMinus, nil)
	case '+':
		tok = sc.instantiateToken(token.TPlus, nil)
	case ';':
		tok = sc.instantiateToken(token.TSemicolon, nil)
	case '*':
		tok = sc.instantiateToken(token.TStar, nil)
	case '!':
		if sc.check('=') {
			tok = sc.instantiateToken(token.TBangEqual, nil)
		} else {
			tok = sc.instantiateToken(token.TBang, nil)
		}

	case '=':
		if sc.check('=') {
			tok = sc.instantiateToken(token.TEqualEqual, nil)
		} else {
			tok = sc.instantiateToken(token.TEqual, nil)
		}

	case '<':
		if sc.check('=') {
			tok = sc.instantiateToken(token.TLessEqual, nil)
		} else {
			tok = sc.instantiateToken(token.TLess, nil)
		}

	case '>':
		if sc.check('=') {
			tok = sc.instantiateToken(token.TGreaterEqual, nil)
		} else {
			tok = sc.instantiateToken(token.TGreater, nil)
		}

	case '/':

		// So we are in comment
		if sc.check('/') {
			for sc.lookahead() != '\n' && sc.hasNextToken() {
				_ = sc.consumeByte()
			}
		} else {
			tok = sc.instantiateToken(token.TSlash, nil)
		}

	case '\n':
		sc.currLine++

	case '"':
		for sc.lookahead() != '"' && sc.hasNextToken() {
			if sc.lookahead() == '\n' {
				sc.currLine++
			}

			sc.consumeByte()
		}

		if !sc.hasNextToken() {
			return sc.instantiateToken(token.TString, "").WithError(
				scanner.ErrLexStringUnterminated,
			)
		}

		sc.consumeByte()
		tok = sc.instantiateToken(token.TString, sc.rawSources[sc.start+1:sc.current-1])

	case ' ', '\t', '\r':
		return nil

	default:

		// So we found a digit literal
		if sc.isDigit(b) {
			// Reading the head
			sc.consumeDigits()

			// Check if there are a fractional part
			if sc.check('.') {
				sc.consumeByte()
				sc.consumeDigits()
			}

			v, err := strconv.ParseFloat(sc.rawSources[sc.start:sc.current], 64)
			if err != nil {
				return sc.instantiateToken(token.TNumber, math.NaN()).WithError(scanner.ErrLexBadNumber)
			}

			tok = sc.instantiateToken(token.TNumber, v)
		} else if sc.isAlphanumeric(b) { // starts with alpha (cause digit is covered upper) | keyword and identifier logic
			for sc.isAlphanumeric(sc.lookahead()) {
				sc.consumeByte()
			}

			literal := sc.rawSources[sc.start:sc.current]
			if tokType, found := token.KeywordsTranslation[literal]; found {
				tok = sc.instantiateToken(tokType, literal)
			} else {
				tok = sc.instantiateToken(token.TIdentifier, literal)
			}
		} else {
			tok = sc.instantiateToken(token.TBadChar, nil).WithError(scanner.ErrLexIncorrectChar)
		}
	}

	return tok
}

func (sc *BaselineScanner) consumeDigits() {
	for sc.isDigit(sc.lookahead()) {
		sc.consumeByte()
	}
}

func (sc *BaselineScanner) isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || sc.isDigit(c)
}

func (sc *BaselineScanner) isDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}

func (sc *BaselineScanner) instantiateToken(tokType token.TokenType, literal any) *token.Token {
	return token.New(tokType, sc.rawSources[sc.start:sc.current], literal, sc.currLine)
}

func (sc *BaselineScanner) consumeByte() byte {
	sc.current++
	return sc.rawSources[sc.current-1]
}

func (sc *BaselineScanner) Tokenize(source string) []*token.Token {
	defer sc.Reset()

	sc.rawSources = source
	tokens := make([]*token.Token, 0, BASE_TOKENS_CAPACITY)

	for sc.hasNextToken() {
		sc.start = sc.current
		if newTok := sc.scanNext(); newTok != nil {
			tokens = append(tokens, newTok)
		}
	}

	// Add EOF
	tokens = append(tokens, token.New(token.TEOF, "", nil, sc.currLine))
	return tokens
}

func (sc *BaselineScanner) Reset() {
	sc.currLine = 1
	sc.start = 0
	sc.current = 0
}

func GetScanner() *BaselineScanner {
	return &BaselineScanner{
		currLine: 1,
		start:    0,
		current:  0,
	}
}
