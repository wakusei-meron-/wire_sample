// +build wireinject
package main

import (
	"github.com/google/wire"

	"wire_sample/controller"
	"wire_sample/env"
	"wire_sample/repository"
	"wire_sample/repository/db"
	"wire_sample/service"
)

func NewController(e env.Conf) controller.Controller {
	wire.Build(
		db.OpenDB,
		repository.NewRepository,
		service.NewService,
		controller.NewController,
	)
	return controller.Controller{}
}
