package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// middleware here
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(MiddlewareOne)
	e.Use(MiddlewareTwo)
	e.Use(echo.WrapMiddleware(MiddlewareSomething))

	e.GET("/index", func(context echo.Context) error {
		fmt.Println(" satu dua tiga")
		return context.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("from middleware one")
		return next(context)
	}
}

func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("from middleware two")
		return next(context)
	}
}

func MiddlewareSomething(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from middleware something")

		h.ServeHTTP(w, r)
	})
}
