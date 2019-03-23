package ast

import "token"

// Node ...
type Node interface {
	TokenLiteral() string
}

// Statement ...
type Statement interface {
	Node
	statementNode()
}

// Expression ...
type Expression interface {
	Node
	expressionNode()
}

// Program ...
type Program struct {
	Statements []Statement
}

// TokenLiteral ...
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement ...
type LetStatement struct {
	Token token.Token // token.Token.LET
	Name  *Identifier
	Value Expression
}

// Identifier ...
type Identifier struct {
	Token token.Token // the token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ...
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
