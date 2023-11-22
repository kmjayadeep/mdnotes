package main

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/kmjayadeep/mdnotes/pkg/handler/render"
)

func main() {
	engine := html.New("./views", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"url": func(s string) template.URL {
			return template.URL(s)
		},
	})
	app := fiber.New(fiber.Config{
		Views:             engine,
		ViewsLayout:       "layouts/main",
		PassLocalsToViews: true,
	})

  app.Static("/assets", "./assets/dist")

  r := render.NewRender()

  app.Get("/", r.Index)

	log.Fatal(app.Listen(":3000"))
}
