package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mycok/monkey_interpreter/lexer"
	"github.com/mycok/monkey_interpreter/token"
)

const prompt = ":: "

// Start displays a user prompt message and initializes a scanner object to read user
// input from the stdOut.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
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
