package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	e := echo.New()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.conf")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	fmt.Println("Starting", viper.GetString("appName"))

	e.GET("/index", func(context echo.Context) error {
		return context.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))

}
