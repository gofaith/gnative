package web

type Promise[T any] struct {
	ID   uint64
	Data T
}
