package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hugoleodev/hug/lexer"
	"github.com/hugoleodev/hug/token"
)

const PROMPT = ">>> "

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

		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Printf("%+v\n", tk)
		}
	}

}
