package web

import "strings"

type ContainerButton struct {
	*Container[*ContainerButton]
}

func Button(texts ...string) *ContainerButton {
	this := &ContainerButton{}
	this.Container = &Container[*ContainerButton]{
		Element: &Element[*ContainerButton]{
			ref:     this,
			TagName: "button",
		},
	}
	if len(texts) > 0 {
		this.Children = append(this.Children, Text(strings.Join(texts, " ")))
	}
	return this
}
