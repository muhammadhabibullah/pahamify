package controller

import (
	"github.com/labstack/echo/v4"

	"pahamify/project/config"
	"pahamify/project/service"
)

// Init REST controller
func Init(svc *service.Container) {
	cfg := config.GetConfig()
	e := echo.New()

	NewPokemonController(e, svc)

	e.Logger.Fatal(e.Start(cfg.Server.Address))
}
