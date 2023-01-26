/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/main.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package main

import (
	"api/cmd/api/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables.
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to read `.env` file.\n%s\n", err)
		os.Exit(1)
	}

	// Creating the API.
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Go-Fiber",
		AppName:       "AISoC API",
	})

	// Registering all routes.
	app.Get("/", routes.IndexRoute)

	// Run!
	app.Listen(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
