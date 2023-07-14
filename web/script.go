package web

type ContainerScript struct {
	*Container[*ContainerScript]
}

func Script(jsCode string) *ContainerScript {
	this := &ContainerScript{}
	this.Container = &Container[*ContainerScript]{
		Element: &Element[*ContainerScript]{
			ref:     this,
			TagName: "script",
		},
	}
	this.Children = append(this.Children, Text(jsCode))
	return this
}

func ScriptSrc(src string) *ContainerScript {
	return Script("").SetAttribute("src", src)
}
