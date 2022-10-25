package checkcustomerror

import (
	"go/ast"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "checkcustomerror is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkcustomerror",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	cmap := getCommentMap(pass)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ReturnStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ReturnStmt:
			last := n.Results[len(n.Results)-1]
			if last == nil {
				return
			}
			lastTyp := pass.TypesInfo.TypeOf(last)
			if lastTyp == nil {
				return
			}

			if lastTyp.String() != "error" {
				return
			}

			ident, _ := last.(*ast.Ident)
			if ident == nil {
				return
			}

			pos := pass.Fset.Position(n.Pos())
			c, ok := cmap[pos.Filename+"_"+strconv.Itoa(pos.Line)]
			if ok {
				if strings.Contains(c, "nocheck:checkcustomerror") {
					return
				}
			}

			pass.Reportf(n.Pos(), "require wrap customError")

		}
	})

	return nil, nil
}

func getCommentMap(pass *analysis.Pass) map[string]string {
	var cmap = make(map[string]string)

	for _, file := range pass.Files {
		for _, cg := range file.Comments {
			for _, c := range cg.List {
				pos := pass.Fset.Position(c.Pos())
				cmap[pos.Filename+"_"+strconv.Itoa(pos.Line)] = c.Text
			}
		}
	}

	return cmap
}
