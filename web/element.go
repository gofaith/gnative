package web

import (
	"fmt"
	"strings"
)

type IElement interface {
	Marshal() string
	MarshalIndent(prefix string) string
}

type Element[T IElement] struct {
	TagName    string
	ref        T // for syntax iteration
	attributes map[string]string
}

func (e *Element[T]) SetAttribute(key string, value string) T {
	if e.attributes == nil {
		e.attributes = make(map[string]string)
	}
	e.attributes[key] = value
	println("SetAttribute: ", key, value)
	return e.ref
}
func (e *Element[T]) GetAttribute(key string) string {
	if e.attributes == nil {
		e.attributes = make(map[string]string)
	}
	return e.attributes[key]
}

func (e *Element[T]) ID(id string) T {
	e.SetAttribute("id", id)
	return e.ref
}

func (e *Element[T]) Marshal() string {
	builder := new(strings.Builder)
	for k := range e.attributes {
		builder.WriteString(" " + k + "=\"" + e.attributes[k] + "\"")
	}
	return fmt.Sprintf("<%s%s>", e.TagName, builder.String())
}

func (e *Element[T]) MarshalIndent(prefix string) string {
	builder := new(strings.Builder)
	for k := range e.attributes {
		builder.WriteString(" " + k + "=\"" + e.attributes[k] + "\"")
	}
	return fmt.Sprintf("%s<%s%s>", prefix, e.TagName, builder.String())
}

func (e *Element[T]) OnClick(a *Application, fn func()) T {
	id := generateMessageId()
	e.SetAttribute("onclick", fmt.Sprintf("_gnativeCallGo(%v)", id))
	a.eventMap[id] = func() {
		fn()
	}
	return e.ref
}
