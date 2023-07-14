package web

import "strings"

type ContainerSpan struct {
	*Container[*ContainerSpan]
}

func Span(texts ...string) *ContainerSpan {
	this := &ContainerSpan{}
	this.Container = &Container[*ContainerSpan]{
		Element: &Element[*ContainerSpan]{
			ref:     this,
			TagName: "span",
		},
	}
	if len(texts) > 0 {
		this.Children = append(this.Children, Text(strings.Join(texts, " ")))
	}
	return this
}
