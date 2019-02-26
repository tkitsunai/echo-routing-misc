package main

import (
	"github.com/tkitsunai/echo-routing-test/router"
)

func main() {
	router.FactoryRouter().Start(":8080")
}
