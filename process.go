package main

import (
	"io"

	blackfriday "github.com/russross/blackfriday/v2"
)

type RenderCode struct{}

func Process(input []byte) ([]byte, error) {
	out := blackfriday.Run(input, blackfriday.WithRenderer(RenderCode{}))
	return out, nil
}

func (_ RenderCode) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) (status blackfriday.WalkStatus) {
	status = blackfriday.GoToNext

	if node.Type != blackfriday.CodeBlock {
		return
	}

	// TODO use the node.CodeBlockData info for situational awareness.
	// Would be nice to avoid dumping code blocks for other languages.

	// Simply dump the contents
	w.Write(node.Literal)
	return
}

func (_ RenderCode) RenderHeader(w io.Writer, ast *blackfriday.Node) {
	// TODO Print a warning that this file was autogenerated, etc.
	// Would be nice to point reader to original file.
	// nop
}

func (_ RenderCode) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	// TODO Print another warning, maybe
	// nop
}