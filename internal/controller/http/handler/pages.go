package handler

import (
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
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

	if err := prepare(e); err != nil {
		log.Error("internal - controller - http - pages.go - NewPages() - prepare(e): %w", err)
		return
	}

	if err := addPages(e); err != nil {
		log.Error("internal - controller - http - pages.go - NewPages() - addPages(e): %w", err)
		return
	}
}

func prepare(e *echo.Echo) error {
	filenames, err := filepath.Abs(filepath.Join("./", "/web/template/*.html"))

	if err != nil {
		return err
	}

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob(filenames)),
	}

	filenames, err = filepath.Abs(filepath.Join("./", "web/static"))

	if err != nil {
		return err
	}

	e.Static("/static/", filenames)

	return nil
}

func addPages(e *echo.Echo) error {
	ph := PagesHandler{}

	e.GET("/", ph.Main)
	e.GET("/signin", ph.Signin)
	e.GET("/signup", ph.Signup)
	e.GET("/wishlist", ph.Wishlist)

	return nil
}

func (h PagesHandler) Main(c echo.Context) error {

	c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title":      "Wish Squishy",
		"authorized": true,
	})

	return nil
}

func (h PagesHandler) Signin(c echo.Context) error {
	c.Render(http.StatusOK, "signin.html", map[string]interface{}{
		"title":      "Wish Squishy",
		"authorized": false,
	})

	return nil
}

func (h PagesHandler) Signup(c echo.Context) error {
	c.Render(http.StatusOK, "signup.html", map[string]interface{}{
		"title":      "Wish Squishy",
		"authorized": false,
	})

	return nil
}

func (h PagesHandler) Wishlist(c echo.Context) error {
	c.Render(http.StatusOK, "wishlist.html", map[string]interface{}{
		"title":      "Wish Squishy",
		"authorized": false,
	})

	return nil
}
