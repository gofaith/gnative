package web

type Window struct {
	app  *Application
	body []IElement
}

func (w *Window) Body(elements ...IElement) *Window {
	w.body = elements
	return w
}

func (w *Window) Export() string {
	return ""
}
