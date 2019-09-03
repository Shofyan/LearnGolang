package main

import (
	"github.com/labstack/echo"
	"github.com/unrolled/secure"
	"net/http"
)

func main() {
	e := echo.New()

	secureMiddleware := secure.New(secure.Options{
		BrowserXssFilter:        true,
		ContentTypeNosniff:      true,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		AllowedHosts:            []string{"localhost:9000", "www.google.com"},
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.GET("/index", func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return context.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))

}
