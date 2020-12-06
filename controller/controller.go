package controller

import (
	"fmt"

	"wire_sample/service"
)

type (
	Controller struct {
		Service service.Service
	}
)

func NewController(service service.Service) Controller {
	return Controller{service}
}

func (c *Controller) Start() {
	fmt.Println("application start")
}
