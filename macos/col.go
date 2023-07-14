package macos

type ContainerCol struct {
	*Container[*ContainerCol]
}

func Col(children ...IElement) *ContainerCol {
	this := &ContainerCol{}
	this.Container = &Container[*ContainerCol]{
		Element: &Element[*ContainerCol]{
			ref:     this,
			TagName: "col",
		},
	}
	if len(children) > 0 {
		this.Children = children
	}
	return this
}
