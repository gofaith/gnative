package web

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type message struct {
	operation string
	id        uint64
	data      []string
}

const (
	operationCall   = "call"
	operationReturn = "return"
)

var (
	messageIdIncrement uint64 = 0
	messageIdLocker    sync.Mutex
)

func generateMessageId() uint64 {
	messageIdLocker.Lock()
	defer messageIdLocker.Unlock()
	messageIdIncrement++
	return messageIdIncrement
}

func parseMessage(s string) (*message, error) {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	var errInvalidMessage = errors.New("invalid message:" + s)
	if len(ss) < 2 {
		return nil, errInvalidMessage
	}

	v := &message{}
	// operation
	var e error
	v.operation, e = url.QueryUnescape(ss[0])
	if e != nil {
		return nil, errInvalidMessage
	}
	switch v.operation {
	case operationCall, operationReturn:
	default:
		return nil, errInvalidMessage
	}

	// id
	v.id, e = strconv.ParseUint(ss[1], 10, 64)
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

func (m *message) Execute(a *Application) {
	switch m.operation {
	case operationCall:
		m.callGo(a)
	}
}
func (m *message) callGo(a *Application) {
	if fn, ok := a.eventMap[m.id]; ok {
		fn()
	} else {
		log.Println("callback of message with id:", m.id, " not found. "+m.String())
	}
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
	return url.QueryEscape(m.operation) + " " + strconv.FormatUint(m.id, 10) + s
}
