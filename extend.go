package links

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var extender *Extender

type Extender struct {
	Source    map[string]bool
	Attribute map[string]string
}

// New return initialized image render with source url replacing support.
func NewExtender(source map[string]bool, attribute map[string]string) goldmark.Extender {
	extender = &Extender{
		Source:    source,
		Attribute: attribute,
	}
	return extender
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewTransformer(), 500),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewRenderer(), 500),
		),
	)
}
