package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
)

// PROMPT ...
const PROMPT = ">>"

// Start ...
func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v \n", tok)
		}
	}

}
