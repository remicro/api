package balancer

import (
	"errors"
	"github.com/remicro/api/app"
	"github.com/remicro/api/cloud/discovery"
)

var (
	ErrUnknownService = errors.New("unknown service")
	ErrAllDeclined = errors.New("all declined")
)

type Balancer interface {
	Find(name string) (node discovery.Node, err error)
	Decline(node discovery.Node) (nextNode discovery.Node, err error)
}

type Builder interface {
	Config(config app.Config) Builder
	Discovery(discoverer discovery.Discovery) Builder
	Build() (balancer Balancer, err error)
}

