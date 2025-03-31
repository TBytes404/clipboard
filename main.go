package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/tbytes404/clipboard/view"
)

var dummy = []string{
	"hello hi goodbye tumi jemon amie temon",
	"hello hi goodbye tumi jemon amie temon",
	"hello hi goodbye tumi jemon amie temon",
	"hello hi goodbye tumi jemon amie temon",
	"hello hi goodbye tumi jemon amie temon",
}

func main() {
	e := echo.New()

	e.Static("/public", "./assets/")

	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, view.HomePage(dummy))
	})

	e.Logger.Fatal(e.Start(":8282"))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}
	return ctx.HTML(statusCode, buf.String())
}
