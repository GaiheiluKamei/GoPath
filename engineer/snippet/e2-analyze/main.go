package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// parse a Go program into an AST representation.
func parse(program string) (*token.FileSet, *ast.File, error) {
	fs := token.NewFileSet()
	tree, err := parser.ParseFile(fs, "example.go", program, 0)
	if err != nil {
		return nil, nil, err
	}

	return fs, tree, nil
}
