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
	"api/pkg/utils"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var API_STATE utils.ENV_VAR_STATE

func init() {
	API_STATE.API_MODE = os.Getenv("API_MODE")
	API_STATE.API_PORT = os.Getenv("API_PORT")
	API_STATE.DATABASE_PATH = os.Getenv("DATABASE_PATH")
	API_STATE.LOGGER = nil
}

func main() {
	// Loading environment variables.
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to read `.env` file.\n%s\n", err)
		os.Exit(1)
	}

	// Since we prefork for performance boost, we need to keep master thread checks!
	if !fiber.IsChild() {
		utils.InitializeLogger(&API_STATE)
		utils.InitializeDbHandle(&API_STATE)
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
