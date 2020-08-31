package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ziwon/dokkery/pkg/app"
	"github.com/ziwon/dokkery/pkg/config"
	"github.com/ziwon/dokkery/pkg/routes"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run App server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	r := echo.New()
	r.HidePort = true
	r.HideBanner = true

	var conf config.Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	app := app.New(r, conf)
	r.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	})
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	routes.Register(r)
	r.Logger.SetLevel(log.DEBUG)

	go func() {
		log.Infof("Starting server on: %v", conf.Server.Address)
		if err := r.Start(conf.Server.Address); err != nil && err != http.ErrServerClosed {
			r.Logger.Error(err)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Errorf("Error on shutdown: %v", err)
	}

	r.Logger.Info("Shutdown")
}
