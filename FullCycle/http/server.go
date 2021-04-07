package http

import (
	"github.com/VictordaSilvaf/FullCycleGoLang/model"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
}
