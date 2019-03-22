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

	MINUS = "-"

	BANG = "!"

	ASTERISK = "*"

	SLASH = "/"

	LT = "<"

	GT = ">"

	EQ = "=="

	NOT_EQ = "!="

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

	TRUE = "TRUE"

	FALSE = "FALSE"

	IF = "IF"

	ELSE = "ELSE"

	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
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
