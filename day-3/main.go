package main

import (
	"github.com/nurhidaylma/alterra-agmc/day-3/config"
	"github.com/nurhidaylma/alterra-agmc/day-3/middlewares"
	"github.com/nurhidaylma/alterra-agmc/day-3/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8080"))
}
