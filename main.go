package main

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "FindHoge",
	Doc:  "Find hoge",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f := n.(*ast.FuncDecl)

		if f.Name.Name == "main" {

			pass.Reportf(f.Pos(), "found main function")

			ast.Inspect(f, func(n ast.Node) bool {
				if n == nil {
					return false
				}

				ast.Print(pass.Fset, n)
				pass.Reportf(n.Pos(), fmt.Sprintf("%#v", n))
				switch n := n.(type) {
				case *ast.CallExpr:
					if sel, ok := n.Fun.(*ast.SelectorExpr); ok {
						if id, ok := sel.X.(*ast.Ident); ok {
							if id.Name == "fmt" && sel.Sel.Name == "Println" {
								pass.Reportf(n.Pos(), "found fmt.Println")
							}
						}
					}

					return false
				}

				return true
			})

			return
		}
	})

	return nil, nil
}

func main() {
	singlechecker.Main(
		Analyzer,
	)
}
