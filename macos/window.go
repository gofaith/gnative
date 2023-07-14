package macos

type ContainerWindow struct {
	*Container[*ContainerWindow]
	app      *Application
	windowID uint64
}

func Window(app *Application, children ...IElement) *ContainerWindow {
	this := &ContainerWindow{
		app:      app,
		windowID: generateRequestId(),
	}
	this.Container = &Container[*ContainerWindow]{
		Element: &Element[*ContainerWindow]{
			ref:     this,
			TagName: "window",
		},
	}
	this.Children = children
	return this
}

func (w *ContainerWindow) Show() *ContainerWindow {
	msg := message{
		method:    methodShow,
		requestID: w.windowID,
		path:      "normal",
		data: []string{
			w.Marshal(),
		},
	}
	msg.ExecuteSwift(w.app)
	return w
}
