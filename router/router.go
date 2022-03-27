package router

import (
	"tf-sso/router/api"
	"tf-sso/service/common"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &common.CustomValidator{Validator: validator.New()}

	apiGroup := e.Group("/api")
	apiGroup.POST("/user", api.CreateUser)
	apiGroup.DELETE("/user/:id", api.DeleteUser)
	apiGroup.GET("/user/:id", api.GetUser)
	apiGroup.GET("/user", api.GetUsers)

	return e
}
