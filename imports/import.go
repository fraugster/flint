package imports

import (
	"context"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/fraugster/flint/plugins"
)

type walker struct {
	err error
	src []byte

	counter int
}

func (fv *walker) Visit(node ast.Node) ast.Visitor {
	if node != nil {
		switch t := node.(type) {
		case *ast.GenDecl:
			fv.counter++
			src := string(fv.src[node.Pos()-1 : node.End()-1])
			if src == `import "C"` { // specifically ignore import "C" as this is a valid case where more than one import block may appear in code.
				fv.counter--
				return fv
			}

			exam := multiLiner
			// One liner import
			if t.Lparen == 0 && t.Rparen == 0 {
				exam = oneLiner
			}

			err := exam(src)
			if err != nil {
				fv.err = err
				return nil
			}
		}
	}

	return fv
}

type checker struct{}

func (checker) Process(ctx context.Context, src []byte) error {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		return err
	}

	w := &walker{
		src: src,
	}

	ast.Walk(w, f)

	if w.counter > 1 {
		return errors.New("two import section in one file, merge them into one")
	}

	return w.err
}

func init() {
	plugins.Register("imports", &checker{})
}
