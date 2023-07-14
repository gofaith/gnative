package macos

type ContainerText struct {
	*Container[*ContainerText]
}

func Text(s string) *ContainerText {
	this := &ContainerText{}
	this.Container = &Container[*ContainerText]{
		Element: &Element[*ContainerText]{
			ref:     this,
			TagName: "text",
		},
	}
	this.Children = append(this.Children, text(s))
	return this
}
