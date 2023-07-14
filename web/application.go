package web

import (
	_ "embed"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

//go:embed ws.js
var wsJs string

type Application struct {
	upgrader           websocket.Upgrader
	WebsocketRoutePath string

	eventMap    map[uint64]func()
	conns       []*websocket.Conn
	connsLocker sync.Mutex
}

func NewApplication(wsRoutePath string) *Application {
	return &Application{
		WebsocketRoutePath: wsRoutePath,
		eventMap:           make(map[uint64]func()),
	}
}

func (a *Application) NewWindow() *Window {
	return &Window{
		app: a,
	}
}
func (a *Application) removeConn(conn *websocket.Conn) {
	a.connsLocker.Lock()
	defer a.connsLocker.Unlock()
	for i, c := range a.conns {
		if c == conn {
			a.conns = append(a.conns[:i], a.conns[i+1:]...)
			return
		}
	}
}
func (a *Application) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	if a.WebsocketRoutePath == "" {
		log.Fatal("Application.WebsocketRoutePath is not set")
		http.Error(w, "WebsocketRoutePath is not set", http.StatusInternalServerError)
		return
	}
	conn, e := a.upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	a.conns = append(a.conns, conn)
	defer a.removeConn(conn)

	// listen message
	for {
		_, m, e := conn.ReadMessage()
		if e != nil {
			log.Println(e)
			break
		}
		msg, e := parseMessage(string(m))
		if e != nil {
			log.Println(e)
			continue
		}
		msg.Execute(a)
	}
}

func (a *Application) WsScriptHandler(w http.ResponseWriter, r *http.Request) {
	if a.WebsocketRoutePath == "" {
		log.Fatal("Application.WebsocketRoutePath is not set")
		http.Error(w, "WebsocketRoutePath is not set", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.Write([]byte(strings.Replace(wsJs, "{{.}}", a.WebsocketRoutePath, -1)))
}

func (a *Application) CallJS(eval string, callback func(any)) {
	msg := &message{
		operation: operationCall,
		id:        generateMessageId(),
		data:      []string{eval},
	}

	for _, c := range a.conns {
		c.WriteMessage(websocket.TextMessage, []byte(msg.String()))
	}
	return
}
