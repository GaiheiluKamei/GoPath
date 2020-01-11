package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GaiheiluKamei/books/sysbook/ch6/advancedPTY/extend/command"
)

func init() {
	command.Register(command.Base{
		Name:   "shuffle",
		Help:   "Shuffle a list of strings",
		Action: shuffleAction,
	})
	command.Register(command.Base{
		Name:   "print",
		Help:   "Prints a file",
		Action: printAction,
	})
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	a := argsScanner{}
	b := bytes.Buffer{}
	fmt.Fprint(w, "** Welcome to PseudoTerm! **\nPlease enter a command.\n")
	for {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get working directory:", err)
			return
		}
		fmt.Fprintf(w, "\n[%s] > ", filepath.Base(pwd))

		a.Reset()
		b.Reset()
		for {
			a.Scan()
			b.Write(s.Bytes())
			extra := a.Parse(&b)
			if extra == "" {
				break
			}
			b.WriteString(extra)
		}
		if command.GetCommand(a[0]).Run(os.Stdin, w, a...) {
			fmt.Fprintln(w)
			return
		}
	}
}

type argsScanner []string

func (a *argsScanner) Reset() { *a = (*a)[0:0] }

func (a *argsScanner) Parse(r io.Reader) (extra string) {
	// s := bufio.NewScanner(r)
	// s.Split(scanargs)
	// for s.Scan() {
	// 	*a = append(*a, s.Text())
	// }
	// if len(*a) == 0 {
	// 	return ""
	// }
	// lastArg := (*a)[len(*a)-1]
	return ""
}
