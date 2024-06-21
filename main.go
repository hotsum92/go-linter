package main

import (
	"go/ast"
	"go/types"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "FindHoge",
	Doc:  "Find hoge",
	Run:  run,
	FactTypes: []analysis.Fact{
		new(HogeFuncFact),
	},
	Requires:   []*analysis.Analyzer{},
	ResultType: reflect.TypeOf(new(HogeFileMap)),
}

type HogeFuncFact struct{}

func (HogeFuncFact) AFact() {}

type HogeFileMap map[*ast.File]struct{}

func run(pass *analysis.Pass) (interface{}, error) {
	var res HogeFileMap = make(map[*ast.File]struct{})
	for _, f := range pass.Files {
		filePath := pass.Fset.File(f.Pos()).Name()
		fileName := filepath.Base(filePath)
		if strings.HasPrefix(fileName, "hoge") {
			res[f] = struct{}{}
		}

		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && strings.HasPrefix(decl.Name.Name, "Hoge") {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					pass.ExportObjectFact(obj, new(HogeFuncFact))
				}
			}
		}
	}
	return &res, nil
}

func main() {
	unitchecker.Main(
		Analyzer,
	)
}
