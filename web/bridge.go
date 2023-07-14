package web

type Bridge interface {
	Call(funcname string, args ...string) Promise[any]
}
