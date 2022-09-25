package main

import (
	"github.com/Rubemlrm/go-gin-sanbox/internal/app"
)

func main() {
	err := app.StartApp()
	if err != nil {
		panic(err)
	}
}
