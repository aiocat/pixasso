// Copyright (C) 2022 aiocat
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

var HCAPTCHA_SECRET string // Hcaptcha secret key variable

func main() {
	godotenv.Load() // Load .env file

	// Start database connection
	err := DATABASE.StartConnection(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Set environment variables
	HCAPTCHA_SECRET = os.Getenv("HCAPTCHA_SECRET")

	engine := html.New("./views", ".html") // Set html engine

	app := fiber.New(fiber.Config{
		Views: engine,
	}) // New fiber app

	// Security check middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		return c.Next()
	})

	app.Static("/", "./static") // Set static folder

	// Setup routes
	app.Post("/api/users", HandlePostUser)
	app.Post("/api/users/auth", HandleAuthUser)
	app.Get("/signin", func(c *fiber.Ctx) error { return c.SendFile("./views/signin.html") })
	app.Get("/signup", func(c *fiber.Ctx) error { return c.SendFile("./views/signup.html") })

	// Listen port
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
