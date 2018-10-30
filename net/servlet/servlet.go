package servlet

import (
	"context"
	"io"
	"net/http"

	"github.com/remicro/api"
	"github.com/remicro/api/cloud/discovery"
	"github.com/remicro/api/logging"
)

type Session interface {
	Set(key string, value interface{})
	Get(key string, value interface{}) (err error)
}

type Headers interface {
	Get(key string) (value string)
	List(key string) (values []string)
	Add(key string, value ...string)
}

type Cookies interface {
	Get(key string) (value http.Cookie)
	Set(key string, value http.Cookie)
}

type Request interface {
	Logger() (logger logging.Logger)

	Context() context.Context
	Cookies() Cookies
	Headers() (headers Headers)
	Session() Session

	Body() (input []byte)
	BodyReader() (input io.Reader)
	BodyObject() interface{}
	Decode(obj interface{}) (err error)

	Close() (err error)
}

type Response interface {
	Logger() (logger logging.Logger)

	Context() context.Context
	Cookies() (cookies Cookies)
	Headers() (headers Headers)
	Session() Session

	BodyWriter() (writer io.Writer)
	SetBody(body []byte)
	SetBodyObject(obj interface{})
	Encode(obj interface{}) (err error)

	Close() (err error)
}

type Handler func(req Request, rsp Response) (err error)

type Middleware func(req Request, rsp Response) (err error)

type Group interface {
	Add(name, route string, handler Handler, middleware ...Middleware) (err error)
}

type Router interface {
	Add(name, route string, handler Handler, middleware ...Middleware) (err error)
	AddGroup(name, route string, group Group, middleware ...Middleware) (err error)
	AddMiddleware(middleware ...Middleware) (err error)
}

type Servlet interface {
	api.Closable
	api.Runnable
}

type Builder interface {
	Name(name string) Builder
	Discovery(discovery discovery.Registry) Builder
	Router(router Router) Builder
	AddRoute(name, route string, handler Handler, middleware ...Middleware) Builder
	AddRouterGroup(name, route string, group Group, middleware ...Middleware) Builder
	AddMiddleware(middleware ...Middleware) Builder
	Build() (slt Servlet, err error)
}
