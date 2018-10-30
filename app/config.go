package app

import "github.com/remicro/api/serialization"

type ZeroValueConstructor func() interface{}

type Watcher interface {
	Int() (ch <- chan int)
	Ints() (ch <- chan []int)
	Int32() (ch <- chan int64)
	Ints32() (ch <- chan []int64)
	Int64() (ch <- chan int64)
	Ints64() (ch <- chan []int64)
	String() (ch <- chan string)
	Strings() (ch <- chan []string)
	Float64(key string) (ch <- chan float64)
	Float32(key string) (ch <- chan float32)
	Object(key string, decoder serialization.Decoder, constructor ZeroValueConstructor) (ch <- chan interface{})
}

type Group interface {
	Int(key string) (value int, err error)
	Ints(key string) (value []int, err error)
	Int32(key string) (value int32, err error)
	Ints32(key string) (value []int32, err error)
	Int64(key string) (value int64, err error)
	Ints64(key string) (value []int64, err error)
	String(key string) (value string, err error)
	Strings(key string) (value []string, err error)
	Float64(key string) (value float64, err error)
	Float32(key string) (value float32, err error)
	Object(key string, decoder serialization.Decoder, object interface{}) error
	Watcher(key string) (watcher Watcher, err error)
}

type Config interface {
	GetGroup(key string) (group Group, err error)
}
