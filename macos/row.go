package macos

type ContainerRow struct {
	*Container[*ContainerRow]
}

func Row(children ...IElement) *ContainerRow {
	this := &ContainerRow{}
	this.Container = &Container[*ContainerRow]{
		Element: &Element[*ContainerRow]{
			ref:     this,
			TagName: "row",
		},
	}
	if len(children) > 0 {
		this.Children = children
	}
	return this
}
