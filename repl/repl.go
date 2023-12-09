package repl

/*
The repl takes in a line of input and passes it through the lexer to produce the AST

Doesn't do much right now, but will eventually take the AST and turn it into instructions
that can be executed.
*/

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Captainmango/monkey/lexer"
	"github.com/Captainmango/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(PROMPT)
		
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}