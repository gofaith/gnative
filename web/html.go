package web

import (
	"fmt"
)

type ContainerHTML struct {
	head, body *Container[*ContainerDiv]
}

func HTML() *ContainerHTML {
	this := &ContainerHTML{}
	this.head = &Container[*ContainerDiv]{
		Element: &Element[*ContainerDiv]{
			TagName: "head",
		},
	}
	this.body = &Container[*ContainerDiv]{
		Element: &Element[*ContainerDiv]{
			TagName: "body",
		},
	}
	return this
}

func (c *ContainerHTML) Head(children ...IElement) *ContainerHTML {
	c.head.Children = children
	return c
}

func (c *ContainerHTML) Body(children ...IElement) *ContainerHTML {
	c.body.Children = children
	return c
}

func (e *ContainerHTML) Marshal() string {
	return fmt.Sprintf("<html>%s%s</html>", e.head.Marshal(), e.body.Marshal())
}

func (e *ContainerHTML) MarshalIndent(prefix string) string {
	return fmt.Sprintf("<html>\n%s\n%s\n</html>", e.head.MarshalIndent("  "), e.body.MarshalIndent("  "))
}
