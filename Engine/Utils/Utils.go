package Utils

import (
	"bytes"
	"errors"
	"os"
	"strings"
)

/*
|=============================================================
|					NewBufferPool()
|=============================================================
|
|
|
|
|=============================================================
*/
func NewBufferPool(InitialCap int, MaxCap int, Factory BufferFactoryMethod, Close BufferCloseMethod) (*Pool[*bytes.Buffer], error) {
	if InitialCap < 0 || MaxCap <= 0 || InitialCap > MaxCap {
		return nil, errors.New("[ ERROR ]: Invalid Capacity Settings")
	}

	PoolShip := &Pool[*bytes.Buffer]{
		Item:    make(chan *bytes.Buffer, MaxCap),
		Factory: Factory,
		Close:   Close,
	}

	for I := 0; I < InitialCap; I++ {
		Client := PoolShip.Factory()
		PoolShip.Item <- Client
	}

	return PoolShip, nil
}

/*
|=============================================================
|					StringInSlice()
|=============================================================
|
|
|
|
|=============================================================
*/
func StringInSlice(Key string, Value []string) bool {
	for _, Ranger := range Value {
		if Ranger == Key {
			return true
		}
	}
	return false
}

/*
|=============================================================
|					IsSystemInTextMode()
|=============================================================
|
|
|
|
|=============================================================
*/
func IsSystemInTextMode() bool {
	for _, Arg := range os.Args {
		if strings.HasPrefix(Arg, "-test.") {
			return true
		}
	}
	return false
}

/*
|=============================================================
|					Get()
|=============================================================
|
|
|
|
|=============================================================
*/
func (p *Pool[T]) Get() T {
	var Item T
	select {
	case Item = <-p.Item:
	default:
		Item = p.Factory()
	}
	return Item
}
