package main

import (
	"github.com/nurhidaylma/alterra-agmc/day-4/config"
	"github.com/nurhidaylma/alterra-agmc/day-4/middlewares"
	"github.com/nurhidaylma/alterra-agmc/day-4/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8080"))
}
