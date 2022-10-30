package checkcustomerror

import (
	"go/ast"
	"go/types"
	"regexp"
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

var (
	excludeRegex string
)

var errType = types.Universe.Lookup("error").Type()

func init() {
	Analyzer.Flags.StringVar(&excludeRegex, "exclude_regex", "_test.go||e2e|generated", "exclude files by regex")
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
			pos := pass.Fset.Position(n.Pos())
			r := regexp.MustCompile(excludeRegex)

			// ファイル名がexcludeRegexにマッチする場合はチェックしない
			if r.MatchString(pos.Filename) {
				return
			}

			if len(n.Results) == 0 {
				return
			}

			last := n.Results[len(n.Results)-1]
			if last == nil {
				return
			}

			lastTyp := pass.TypesInfo.TypeOf(last)
			if lastTyp != errType {
				return
			}

			if !check(pass, last) {
				return
			}

			c, ok := cmap[pos.Filename+"_"+strconv.Itoa(pos.Line)]
			if ok {
				if strings.Contains(c, "nocheck:checkcustomerror") {
					return
				}
			}

			pass.Reportf(n.Pos(), "require customError wrap")

		}
	})

	return nil, nil
}

func check(pass *analysis.Pass, item ast.Expr) bool {
	ident, _ := item.(*ast.Ident)
	if ident != nil {
		return true
	}

	call := item.(*ast.CallExpr)
	if call == nil {
		return false
	}
	obj := getFun(pass, call.Fun)
	if obj == nil {
		return false
	}

	// メソッド名がCustomErrorでない場合はチェック
	return obj.Name() != "CustomError"
}

func getFun(pass *analysis.Pass, fun ast.Expr) *types.Func {
	switch fun := fun.(type) {
	case *ast.Ident:
		obj, _ := pass.TypesInfo.ObjectOf(fun).(*types.Func)
		return obj
	case *ast.SelectorExpr:
		obj, _ := pass.TypesInfo.ObjectOf(fun.Sel).(*types.Func)
		return obj
	}
	return nil
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
