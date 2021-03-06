package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"unicode"
	"unicode/utf8"

	"github.com/agnivade/levenshtein"
)

type cmd struct {
	Name   string
	Help   string
	Action func(w io.Writer, args ...string) bool
}

func (c cmd) Match(s string) bool {
	return c.Name == s
}

func (c cmd) Run(w io.Writer, args ...string) bool {
	return c.Action(w, args...)
}

var cmds = make([]cmd, 0, 10)

func init() {
	cmds = append(cmds,
		cmd{
			Name: "exit",
			Help: "Exits the application",
			Action: func(w io.Writer, args ...string) bool {
				fmt.Fprintf(w, "Goodbye! :)\n")
				return true
			},
		},
		cmd{
			Name: "help",
			Help: "Shows available commands",
			Action: func(w io.Writer, args ...string) bool {
				fmt.Fprintln(w, "Available Commands:")
				for _, c := range cmds {
					// "%-15s": "-" represent align left, "15" represent the length of character.
					fmt.Fprintf(w, " - %-15s %s\n", c.Name, c.Help)
				}
				return false
			},
		},
		cmd{
			Name: "shuffle",
			Help: "shuffle a list of strings",
			Action: func(w io.Writer, args ...string) bool {
				rand.Shuffle(len(args), func(i, j int) {
					args[i], args[j] = args[j], args[i]
				})
				for i := range args {
					if i > 0 {
						fmt.Fprint(w, " ")
					}
					fmt.Fprintf(w, "%s", args[i])
				}
				fmt.Fprintln(w)
				return false
			},
		},
		cmd{
			Name: "print",
			Help: "Print a file",
			Action: func(w io.Writer, args ...string) bool {
				if len(args) != 1 {
					fmt.Fprintf(w, "Please specify one file!")
					return false
				}
				f, err := os.Open(args[0])
				if err != nil {
					fmt.Fprintf(w, "Cannot open %s: %s\n", args[0], err)
				}
				defer f.Close()
				if _, err := io.Copy(w, f); err != nil {
					fmt.Fprintf(w, "Cannot print %s: %s\n", args[0], err)
				}
				fmt.Fprintln(w)
				return false
			},
		},
	)
}

type argsScanner []string

func (a *argsScanner) Reset() {
	*a = (*a)[0:0]
}

// This custom slice will allow us to receive lines with quotes and new lines between quotes
// by changing how the look works.
func (a *argsScanner) Parse(r io.Reader) (extra string) {
	s := bufio.NewScanner(r)
	s.Split(ScanArgs)
	for s.Scan() {
		*a = append(*a, s.Text())
	}
	if len(*a) == 0 {
		return ""
	}
	lastArg := (*a)[len(*a)-1]
	if !isQuote(rune(lastArg[0])) {
		return ""
	}
	*a = (*a)[:len(*a)-1]
	return lastArg + "\n"
}

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}

/*
ScanArgs is a custom split function that behaves like `bufio.ScanWords` apart from the
fact that it's aware of quotes.

The function has a first block that skips spaces and finds the first non-space character;
if that character is a quote, it is skipped. Then it looks for the first character that
terminates the arguments, which is a space for normal arguments, and the respective quote
for the other arguments.

If the end of file is reached while in a quoted context, the partial string is returned;
otherwise, the quote is not skipped and more data is requested.
*/
func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// first space
	start, first := 0, rune(0)
	for width := 0; start < len(data); start += width {
		first, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(first) {
			break
		}
	}
	// skip quote
	if isQuote(first) {
		start++
	}
	// loop until arg end character
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if ok := isQuote(first); !ok && unicode.IsSpace(r) || ok && r == first {
			return i + width, data[start:i], nil
		}
	}

	// token from EOF
	if atEOF && len(data) > start {
		if isQuote(first) {
			start--
		}
		return len(data), data[start:], nil
	}
	if isQuote(first) {
		start--
	}
	return start, nil, nil
}

func commandNotFound(w io.Writer, cmd string) {
	var list []string
	for _, c := range cmds {
		d := levenshtein.ComputeDistance(c.Name, cmd)
		if d < 3 {
			list = append(list, c.Name)
		}
	}
	fmt.Fprintf(w, "Command %q not found.", cmd)
	if len(list) == 0 {
		return
	}
	fmt.Fprint(w, " Maybe you meant: ")
	for i := range list {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%s", list[i])
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	a := argsScanner{}
	b := bytes.Buffer{}
	fmt.Fprint(w, "** Welcome to pseudoTerm! **\nPlease enter a command.\n")
	for {
		// prompt message
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get working directory:", err)
			return
		}
		fmt.Fprintf(w, "\n[%s] > ", filepath.Base(pwd))

		a.Reset()
		b.Reset()
		for {
			s.Scan()
			b.Write(s.Bytes())
			extra := a.Parse(&b)
			if extra == "" {
				break
			}
			b.WriteString(extra)
		}
		// a contains the split arguments
		idx := -1
		for i := range cmds {
			if !cmds[i].Match(a[0]) {
				continue
			}
			idx = i
			break
		}
		if idx == -1 {
			commandNotFound(w, a[0])
			continue
		}
		if cmds[idx].Run(w, a[1:]...) {
			fmt.Fprintln(w)
			return
		}
	}
}
