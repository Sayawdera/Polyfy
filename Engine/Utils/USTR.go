package Utils

import "bytes"

/*
|=============================================================
|					Pool
|=============================================================
|
|
|
|
|=============================================================
*/
type Pool[T any] struct {
	Item     chan T
	Factory  func() T
	Close    func(T)
	AfterPut func(T)
}

type BufferFactoryMethod func() *bytes.Buffer
type BufferCloseMethod func(*bytes.Buffer)
