package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tkitsunai/echo-routing-test/router/handler"
)

type Router interface {
	Start(address string) error
}

func FactoryRouter() Router {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userIdResource, handle := handler.UsersId()
	e.GET(userIdResource, handle)

	userIdOverviewResource, shandle := handler.UserIdOverviews()
	e.GET(userIdOverviewResource, shandle)

	return e
}
