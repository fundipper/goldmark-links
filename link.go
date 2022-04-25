package links

import (
	"github.com/yuin/goldmark/ast"
)

// KindLink is a NodeKind of the link node.
var KindLink = ast.NewNodeKind("Link")

// Link struct represents a link of the Markdown text.
type Link struct {
	ast.Link
}

// NewLink returns a new link node.
func NewLink(link *ast.Link) *Link {
	c := &Link{
		Link: *link,
	}

	c.Destination = link.Destination
	c.Title = link.Title
	return c
}

// Kind implements Node.Kind.
func (l *Link) Kind() ast.NodeKind {
	return KindLink
}
