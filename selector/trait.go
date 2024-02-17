package selector

type trait[T any] struct {
	Active bool
	Data   T
}
