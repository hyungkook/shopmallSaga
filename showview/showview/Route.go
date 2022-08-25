package showview

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	shopmall := &Shopmall{}
	e.GET("/shopmalls", shopmall.Get)
	e.GET("/shopmalls/:id", shopmall.GetbyId)
	e.POST("/shopmalls", shopmall.Persist)
	e.PUT("/shopmalls/:id", shopmall.Put)
	e.DELETE("/shopmalls/:id", shopmall.Remove)
	return e
}
