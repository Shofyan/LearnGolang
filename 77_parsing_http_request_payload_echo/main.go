package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
		// JSON
		//  curl -X POST http://localhost:9000/user -d "name=Joe" -d "email=nope@novalagung.com"
		// XML
		//curl -X POST http://localhost:9000/user -H "Content-Type: application/xml" -d "<?xml version="1.0"?> <Data> <Name>Joe</Name> <Email>nope@novalagung.com</Email></Data>"
		// query string
		//curl -X GET "http://localhost:9000/user?name=Joe&email=nope@novalagung.com"
	})

	fmt.Println("server started a localhost:9000")
	r.Start(":9000")
}
