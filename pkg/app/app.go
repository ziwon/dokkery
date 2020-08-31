package app

import (
	"github.com/labstack/echo/v4"
	"github.com/ziwon/dokkery/pkg/config"
)

type App struct {
	R      *echo.Echo
	Config config.Config
}

func New(r *echo.Echo, config config.Config) *App {
	return &App{
		R:      r,
		Config: config,
	}
}
