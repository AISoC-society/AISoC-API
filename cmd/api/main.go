/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/main.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package main

import (
	"api/cmd/api/router"
	"api/pkg/utils"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var API_STATE utils.APP_STATE

func init() {
	utils.IsNotFiberChild(func() {
		utils.LoadEnvVars(&API_STATE)
		utils.InitializeLogger(&API_STATE)
		API_STATE.LOGGER.Info("Loaded environment variables from `.env` file successfully: ", API_STATE)

		utils.InitializeDbHandle(&API_STATE)
	})
}

func main() {
	// Creating the API.
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Go-Fiber",
		AppName:       "AISoC API",
	})

	// Registering all routes.
	router.SetupRoutes(app, &API_STATE)

	// Run!
	app.Listen(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
