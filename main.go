package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(filename string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))

	if err != nil {
		return err
	}

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}
	return nil
}

func main() {
	Parse("sample/main.go")
}
