package token

// TokenType ...
type TokenType string

const (
	// ILLEGAL ... and EOF
	ILLEGAL = "ILLEGAL"

	// EOF ...
	EOF = "EOF"

	// IDENT ...  Identifiers
	IDENT = "IDENT" // add, foobar, x, y

	// INT ... literals
	INT = "INT" // 12345

	// ASSIGN ... and PLUS operators
	ASSIGN = "="

	// PLUS ...
	PLUS = "+"

	// COMMA ... delimiters
	COMMA = ","

	// SEMICOLON ...
	SEMICOLON = ";"

	// LPAREN ...
	LPAREN = "("

	// RPAREN ...
	RPAREN = ")"

	// LBRACE ...
	LBRACE = "{"

	// RBRACE ...
	RBRACE = "}"

	// FUNCTION ...  and LET keywords
	FUNCTION = "FUNCTION"

	// LET ...
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Token ...
type Token struct {
	Type    TokenType
	Literal string
}

// LookupIdent ...
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
