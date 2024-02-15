package colleselector

type Trait[T any] struct {
	Active bool
	Data   T
}

type Contains struct {
	Contains      string
	CaseSensitive bool
}

type Parent struct {
	Selector string
	Rootwise bool
}
