package macos

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Application struct {
	bridge      IBridge
	windows     []*ContainerWindow
	upgrader    websocket.Upgrader
	eventMap    map[uint64]func()
	conns       []*websocket.Conn
	connsLocker sync.Mutex
}

func NewApplication(bridge IBridge) *Application {
	return &Application{
		bridge:   bridge,
		eventMap: make(map[uint64]func()),
	}
}
func (a *Application) ExecuteGo(s string) error {
	m, e := parseMessage(s)
	if e != nil {
		return e
	}
	m.ExecuteGo(a)
	return nil
}
func (a *Application) NewWindow() *ContainerWindow {
	w := Window(a)
	a.windows = append(a.windows, w)
	return w
}
