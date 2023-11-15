package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENT = "IDENT" // the name of a function or variable; foo, x, add
	INT   = "INT"   // any whole real number

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
