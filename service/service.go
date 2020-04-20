package service

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/kubernetes"

	"github.com/micro/go-plugins/client/selector/static"
)

func NewK8sService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}

func NewNormalService(opts ...micro.Option) micro.Service {
	return micro.NewService(
		micro.Name("auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
}
