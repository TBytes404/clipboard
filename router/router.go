package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tbytes404/clipboard/handler"
)

func Init(bh *handler.BlobsHandler) *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.Logger.Error(err)
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Static("/public", "./assets/")

	e.GET("/", bh.HomePage)

	e.POST("/", bh.AddBlob)

	e.DELETE("/", bh.DeleteBlob)

	e.GET("/blobs", bh.FetchBlobs)

	return e
}
