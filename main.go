package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	port := flag.String("port", "8080", "port to use")
	devMode := flag.Bool("dev", false, "enable dev mode")
	flag.Parse()

	templateEngine := html.New("./views", ".html")
	templateEngine.Reload(*devMode)

	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})

	app.Use(logger.New())

	setupRouting(app)

	log.Println(log.Ldate|log.Ltime|log.Lshortfile, fmt.Sprintf("Started service v1 on port: : %s. Dev mode is: %t", *port, *devMode))
	log.Fatal(app.Listen(":" + *port))
}

func setupRouting(app *fiber.App) {
	// => http://localhost:3000/static/js/script.js
	// => http://localhost:3000/static/css/style.css
	app.Static("/static", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", getContent("index"))
	})
}

func getContent(name string) *fiber.Map {
	jsonFile, err := os.Open(fmt.Sprintf("content/%s.json", name))

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result fiber.Map
	if err := json.Unmarshal(byteValue, &result); err != nil {
		fmt.Println(err)
	}

	return &result
}
