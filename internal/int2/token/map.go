package token

func init() {
	for k, v := range KeywordsTranslation {
		revertedTranslation[v] = k
	}
}

var KeywordsTranslation = map[string] TokenType {
	"and": TAnd,
	"class": TClass,
	"else": TElse,
	"false": TFalse,
	"fun": TFun,
	"for": TFor,
	"if": TIf,
	"nil": TNil,
	"or": TOr,
	"print": TPrint,
	"return": TReturn,
	"super": TSuper,
	"this": TThis,
	"true": TTrue,
	"var": TVar,
	"while": TWhile,
}

var revertedTranslation = make(map[TokenType] string, 0)