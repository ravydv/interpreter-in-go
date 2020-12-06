package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ravydv/interpreter-in-go/lexer"
	"github.com/ravydv/interpreter-in-go/token"
)

const PROMPT = ">> "

// Start Read, Eval, Print, Loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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
