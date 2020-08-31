package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ziwon/dokkery/pkg/handlers"
)

func Register(r *echo.Echo) {
	r.GET("/", handleIndexPage)

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/event", handlers.HandleEventPullOrPush)
}

func handleIndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
