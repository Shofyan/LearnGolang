package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func MakeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report.Message = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}

func main() {
	e := echo.New()

	// middleware here
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(MiddlewareLogging)
	e.HTTPErrorHandler = ErrorHandler

	e.Use(MiddlewareOne)
	e.Use(MiddlewareTwo)
	e.Use(echo.WrapMiddleware(MiddlewareSomething))

	e.GET("/index", func(context echo.Context) error {
		fmt.Println(" satu dua tiga")
		return context.JSON(http.StatusOK, true)
	})

	//e.Logger.Fatal(e.Start(":9000"))

	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":9000") }(lock)

	time.Sleep(1 * time.Second)
	MakeLogEntry(nil).Warning("Application start with ssl/tls enable ")

	err := <-lock
	if err != nil {
		MakeLogEntry(nil).Panic("failed start application")
	}

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
