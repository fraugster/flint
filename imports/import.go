package imports

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/fraugster/flint/plugins"
	"github.com/pkg/errors"
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
