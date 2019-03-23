package parser

import (
	"ast"
	"lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `

let x = 5;
let y = 10;
let foobar = 838383;
`
	lex := lexer.New(input)
	p := New(lex)
	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	return true
}
