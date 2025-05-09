package main

import (
	"log"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
)

func main(){
html_engine:= html.New("./views",".html")

app:=fiber.New(
	fiber.Config{
		Views: html_engine,
	})
	

	app.Static("/static","./static")

	app.Get("/",func(c *fiber.Ctx)error{
		return c.Render("index",fiber.Map{
			"Title": "Welcome to Fiber",
		})
	})


	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("This is the About Page")
	})

	log.Fatal(app.Listen(":5030"))



}
