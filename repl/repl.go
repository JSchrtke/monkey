package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Fprintf(output, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line);

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(output, "%+v\n", t);
		}
	}
}
