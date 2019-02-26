package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"path"
)

const usersIdBasePath = "/users/:id"

func UsersId() (string, echo.HandlerFunc) {
	return usersIdBasePath, func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("user id: %v", c.Param("id")))
	}
}

func UserIdOverviews() (string, echo.HandlerFunc) {
	return path.Join(usersIdBasePath, "overviews"), func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("user id overviews: %v", c.Param("id")))
	}
}
