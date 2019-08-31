package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

func main() {
	//kingpin.Parse()

	app := kingpin.New("app", "sipmle app")
	argsAppName := kingpin.Arg("name", "application name").Required().String()
	argsPort := kingpin.Arg("port", "web server port").Default("9000").Int()

	_, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error command")
	}

	appName := *argsAppName
	port := fmt.Sprintf(":%d", *argsPort)

	fmt.Println("starting %s at %s", appName, port)
	e := echo.New()
	e.GET("/index", func(context echo.Context) (err error) {
		return context.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(port))

}
