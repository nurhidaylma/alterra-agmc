package main

import (
	"github.com/nurhidaylma/alterra-agmc/day-2/config"
	"github.com/nurhidaylma/alterra-agmc/day-2/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
