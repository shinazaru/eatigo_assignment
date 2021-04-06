package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log.Println("Service API FOR Shortly Domain Starting")

	SetupMongoCollectionIndex()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.POST("/", CreateShortlyHandler)
	e.GET("/:id", RedirectShortlyHandler)
	e.DELETE("/:id", DeleteShortlyHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
