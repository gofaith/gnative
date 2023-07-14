package macos

import (
	"fmt"
	"strings"
)

type Container[T IElement] struct {
	*Element[T]
	Children []IElement
}

func (e *Container[T]) Body(children ...IElement) T {
	e.Children = children
	return e.ref
}

func (e *Container[T]) Marshal() string {
	attr := new(strings.Builder)
	for k := range e.attributes {
		v := e.attributes[k]
		attr.WriteString(" " + k + "=\"" + v + "\"")
	}
	body := new(strings.Builder)
	for _, v := range e.Children {
		body.WriteString(v.Marshal())
	}
	return fmt.Sprintf("<%s%s>%s</%s>", e.TagName, attr.String(), body.String(), e.TagName)
}

func (e *Container[T]) MarshalIndent(prefix string) string {
	attr := new(strings.Builder)
	for k := range e.attributes {
		v := e.attributes[k]
		attr.WriteString(" " + k + "=\"" + v + "\"")
	}
	body := new(strings.Builder)
	for _, v := range e.Children {
		body.WriteString(v.MarshalIndent(prefix+"  ") + "\n")
	}
	return fmt.Sprintf("%s<%s%s>\n%s%s</%s>", prefix, e.TagName, attr.String(), body.String(), prefix, e.TagName)
}
