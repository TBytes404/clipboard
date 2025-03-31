package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/tbytes404/clipboard/db"
	"github.com/tbytes404/clipboard/handler"
	"github.com/tbytes404/clipboard/store"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewBlobsHandler(store.NewBlobsStore(db))

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.Logger.Error(err)
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Static("/public", "./assets/")

	e.GET("/", h.HomePage)

	e.POST("/", h.AddBlob)

	e.DELETE("/", h.DeleteBlob)

	e.Logger.Fatal(e.Start(":8282"))
}
