package ioc

type Container interface {
	Provide(fn interface{})
	Invoke(fn interface{})
}
