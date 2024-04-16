package main

import (
	. "github.com/CarlKlagba/gameplan/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
	"net/http"
	"text/template"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, paths ...string) {
	tmpl := &template.Template{}
	for i := range paths {
		template.Must(tmpl.ParseGlob(paths[i]))
	}
	t := newTemplate(tmpl)
	e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template{
		Templates: templates,
	}
}

var gameplans = map[int]*Gameplan{}

func main() {
	// Create a new echo instance
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	NewTemplateRenderer(e, "public/*.html")

	// Routes
	e.GET("/", mainPage)
	e.GET("/gameplans", getGameplans)
	e.POST("/action/reaction", createReaction)
	e.POST("/gameplans", createGameplan)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func mainPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)

}

func createReaction(c echo.Context) error {
	log.Println("create reaction")
	return c.Render(
		http.StatusOK,
		"reaction",
		Reaction{Name: "TEST Reaction"})
}

func createGameplan(c echo.Context) error {
	g := new(Gameplan)
	if err := c.Bind(g); err != nil {
		return err
	}
	gameplans[len(gameplans)+1] = g
	return c.JSON(200, g)
}

func getGameplans(c echo.Context) error {
	var gameplan1 = NewGameplan("passing to north south")
	var action1 = gameplan1.CreateFirstAction("get angle")
	reaction1 := action1.AddReaction("use bottom leg")
	reaction1.AddAction("Run back to the legs")
	reaction2 := action1.AddReaction("use top leg")
	reaction2.AddAction("Limp arm out and go to North South")
	gameplans[1] = gameplan1

	log.Println(gameplan1)
	log.Println(gameplan1.Action)
	return c.Render(http.StatusOK, "game_plan", gameplan1)
}
