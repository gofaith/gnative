package web

type ContainerDiv struct {
	*Container[*ContainerDiv]
}

func Div(children ...IElement) *ContainerDiv {
	this := &ContainerDiv{}
	this.Container = &Container[*ContainerDiv]{
		Element: &Element[*ContainerDiv]{
			ref:     this,
			TagName: "div",
		},
	}
	if len(children) > 0 {
		this.Children = children
	}
	return this
}

func DivAttr() *ContainerDiv {
	this := &ContainerDiv{}
	this.Container = &Container[*ContainerDiv]{
		Element: &Element[*ContainerDiv]{
			ref:     this,
			TagName: "div",
		},
	}
	return this
}
