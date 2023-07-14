package src

import (
	"github.com/gofaith/gnative/macos"
)

type IBridge interface {
	macos.IBridge
}
type Application struct {
	*macos.Application
}

func (a *Application) ExecuteGo(s string) error {
	return a.Application.ExecuteGo(s)
}

func Main(bridge IBridge) *Application {
	a := macos.NewApplication(bridge)
	w := a.NewWindow()
	w.Body(
		macos.Col(
			macos.Text("asd"),
			macos.Text("pp"),
			macos.Row(
				macos.Text("two"),
				macos.Text("trhee"),
			),
		),
	).Show()

	return &Application{Application: a}
}
