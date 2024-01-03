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

	"github.com/Captainmango/monkey/evaluator"
	"github.com/Captainmango/monkey/lexer"
	"github.com/Captainmango/monkey/object"
	"github.com/Captainmango/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	env := object.NewEnvironment()

	for {
		fmt.Println(PROMPT)
		
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluatedProgram := evaluator.Eval(program, env)

		if evaluatedProgram != nil {
			io.WriteString(out, evaluatedProgram.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Oh dear. Our monkies saw your code and it doesn't look quite right...\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
