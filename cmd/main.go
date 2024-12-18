package main

import (
	"github.com/VeerDan/calc_go/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}