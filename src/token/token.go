package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"     
	EOF     = "EOF"    

	// Identifiers + literals	
	//
	IDENT = "IDENT" // add, foobar, x, y
	INT = "INT"		// 12345

	// Operators 
	//
	ASSIGN = "="
	PLUS = "+"
	
	// Delimiters
	//
	COMMA = ","
	SEMICOLON = ";"


	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	//
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type Token struct {
	Type TokenType
	Literal string
}