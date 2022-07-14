package http

import (
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/config"
	log "github.com/reucot/wish-squish/pkg/logger"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

type PagesHandler struct {
}

func NewPages(e *echo.Echo) {

	if err := initRender(e); err != nil {
		log.Error("internal - controller - http - pages.go - NewPages(): %w", err)
		return
	}

	ph := PagesHandler{}

	e.GET("/", ph.Main)
	e.GET("/signin", ph.Signin)

}

func initRender(e *echo.Echo) error {
	filenames, err := filepath.Abs(filepath.Join("./", "/web/template/*.html"))

	if err != nil {
		return err
	}

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob(filenames)),
	}

	return nil
}

func (h PagesHandler) Main(c echo.Context) error {

	c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title": "Index",
	})

	return nil
}

func (h PagesHandler) Signin(c echo.Context) error {
	c.Render(http.StatusOK, "signin.html", map[string]interface{}{
		"Host":  config.Get().Host,
		"Title": "Index",
	})

	return nil
}
