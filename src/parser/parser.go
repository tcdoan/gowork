package parser

import (
	"ast"
	"lexer"
	"token"
)

// Parser ...
type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New ...
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

// ParseProgram ...
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
