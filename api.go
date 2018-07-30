package api

type Closable interface {
	Close() (err error)
}

type Runnable interface {
	Run() (err error)
}
