package discovery

import "github.com/remicro/api/app"

type Node interface {
	ID() string
	Address() string
	Schema() string
	Name() string
	Version() string
	Options() ([]Option)
}

type Option interface {
	Type() int
	Value() interface{}
}

type Discovery interface {
	SearchByName(name string) (nodes []Node, err error)
	SearchByID(ID string) (node Node, err error)
}

type Registry interface {
	Retain(node Node, options...Option) (ID string, err error)
	Forget(ID string) (err error)
}

type Factory interface {
	GetDiscovery() Discovery
	GetRegistry() Registry
}

type Builder interface {
	Config(config app.Config) Builder
	Build() Factory
}