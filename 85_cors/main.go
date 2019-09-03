package main

import (
	"github.com/labstack/echo"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	e := echo.New()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://novalagung.com", "https://www.google.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {

		//	writer.Header().Set("Access-Control-Allow-Origin","https://www.google.com, https://novalagung.com")
		//	writer.Header().Set("Access-Control-Allow-Origin","*")
		//	writer.Header().Set("Access-Control-Allow-Methods","OPTIONS, GET, POST, PUT")
		//	writer.Header().Set("Access-Control-Allow-Headers","Content-Type, X-CSRF-Token")

		//if request.Method == "OPTIONS"{
		//	writer.Write([]byte("allowed"))
		//	return
		//}
		writer.Write([]byte("hello"))
	})

	//log.Println("Starting app at :9000")
	//http.ListenAndServe(":9000",nil)

	e.Logger.Fatal(e.Start(":9000"))

}
