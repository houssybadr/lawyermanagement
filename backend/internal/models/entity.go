package models

type Entity interface {
	ToString() string
	GetId() uint32
}

type Jsonable[T any] interface {
	FromJson() error
	ToJson(t T) (string, error)
}

type Copyable[T any] interface {
	CopyWith(T)
}
