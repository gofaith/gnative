package macos

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type message struct {
	method    string
	requestID uint64
	path      string
	data      []string
}

const (
	methodCall   = "call"
	methodReturn = "return"
	methodShow   = "show" // show new window
)

var (
	requestIdIncrement uint64 = 0
	requestIdLocker    sync.Mutex
)

func generateRequestId() uint64 {
	requestIdLocker.Lock()
	defer requestIdLocker.Unlock()
	requestIdIncrement++
	return requestIdIncrement
}

func parseMessage(s string) (*message, error) {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	var errInvalidMessage = errors.New("invalid message:" + s)
	if len(ss) < 3 {
		return nil, errInvalidMessage
	}

	v := &message{}
	// operation
	var e error
	v.method, e = url.QueryUnescape(ss[0])
	if e != nil {
		return nil, errInvalidMessage
	}
	switch v.method {
	case methodCall, methodReturn:
	default:
		return nil, errInvalidMessage
	}

	// id
	v.requestID, e = strconv.ParseUint(ss[1], 10, 64)
	if e != nil {
		return nil, errInvalidMessage
	}

	// path
	v.path, e = url.QueryUnescape(ss[2])
	if e != nil {
		return nil, errInvalidMessage
	}

	for i := 3; i < len(ss); i++ {
		arg, e := url.QueryUnescape(ss[i])
		if e != nil {
			return nil, errInvalidMessage
		}

		v.data = append(v.data, arg)
	}
	return v, nil
}

func (m *message) ExecuteGo(a *Application) {
	switch m.method {
	case methodCall:
		if fn, ok := a.eventMap[m.requestID]; ok {
			fn()
		} else {
			log.Println("callback of message with id:", m.requestID, " not found. "+m.String())
		}
	}
}

func (m *message) ExecuteSwift(a *Application) {
	a.bridge.ExecuteSwift(m.String())
}

func (m *message) String() string {
	s := ""
	if len(m.data) > 0 {
		s = " "
		for i, v := range m.data {
			s += url.QueryEscape(v)
			if i < len(m.data)-1 {
				s += " "
			}
		}
	}
	return url.QueryEscape(m.method) + " " + strconv.FormatUint(m.requestID, 10) + " " + url.QueryEscape(m.path) + s
}
