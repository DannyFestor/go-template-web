package controllers

import (
	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

type Controllers struct {
	ErrorController *ErrorController
	HomeController  *HomeController
	UserController  *UserController
}

func Init(app *config.Application) *Controllers {
	return &Controllers{
		ErrorController: &ErrorController{app: app},
		HomeController:  &HomeController{app: app},
		UserController:  &UserController{app: app},
	}
}
