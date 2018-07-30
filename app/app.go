package app

import (
	"github.com/remicro/api"
	"github.com/remicro/api/ioc"
	"context"
)

type Application interface {
	api.Closable
	api.Runnable
}

type Builder interface {
	Container(container ioc.Container) (Builder)
	Name(name string) (Builder)
	Version(name string) (Builder)
	Context(ctx context.Context) (Builder)

	Build() (app Application)
}