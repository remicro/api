package ioc

type Container interface {
	Provide(fn interface{}) Container
	Invoke(fn interface{}) Container
	Inject() (err error)
}
