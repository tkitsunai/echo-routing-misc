package main

import (
	"github.com/tkitsunai/echo-routing-misc/router"
)

func main() {
	router.FactoryRouter().Start(":8080")
}
