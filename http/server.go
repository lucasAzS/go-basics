package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasAzS/legendary-computing-machine/model"
)
type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer{
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context)	error {
	return c.JSON(http.StatusOK, w.Products)
}