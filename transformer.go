package links

import (
	"fmt"
	"net/url"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (a *Transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if n.Kind() != ast.KindLink {
			return ast.WalkContinue, nil
		}

		link := n.(*ast.Link)
		u, err := url.Parse(string(link.Destination))
		if err != nil {
			msg := ast.NewString([]byte(fmt.Sprintf("<!-- %s -->", err)))
			msg.SetCode(true)
			node.Parent().InsertAfter(node.Parent(), node, msg)
			return ast.WalkContinue, nil
		}

		if u.Host == "" {
			return ast.WalkContinue, nil
		}

		state, ok := extender.Source[u.Host]
		if ok && state {
			return ast.WalkContinue, nil
		}

		n.Parent().ReplaceChild(n.Parent(), n, NewLink(link))

		return ast.WalkContinue, nil
	})
}
