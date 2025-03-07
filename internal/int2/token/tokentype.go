package token

type TokenType uint8

const (
	// Single-char tokens
	TLParent TokenType = iota 
	TRParent 				
	TLBrace
	TRBrace
	TComma
	TDot
	TMinus
	TPlus
	TSemicolon
	TSlash
	TStar
	
	// 1-2 chars
	TBang
	TBangEqual
	TEqual
	TEqualEqual
	TGreater
	TGreaterEqual
	TLess
	TLessEqual
	
	// Literals
	TIdentifier
	TString
	TNumber
	
	// Keywords
	TAnd
	TClass
	TElse
	TFalse
	TFun
	TFor
	TIf
	TNil
	TOr
	TPrint
	TReturn
	TSuper
	TThis
	TTrue
	TVar
	TWhile
	
	TEOF
	TBadChar
)

