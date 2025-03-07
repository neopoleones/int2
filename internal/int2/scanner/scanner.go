package scanner

import "github.com/destr4ct/int2/internal/int2/token"

type Scanner interface {
	Tokenize(source string) []*token.Token
	Reset()
}
