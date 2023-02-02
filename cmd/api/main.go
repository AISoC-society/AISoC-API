/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/main.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package main

import (
	"fmt"

	"github.com/aisoc-society/aisoc-api/cmd/api/router"
	"github.com/aisoc-society/aisoc-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var API_STATE utils.APP_STATE

func init() {
	utils.IsNotFiberChild(func() {
		API_STATE.LoadEnvVars()
		API_STATE.InitializeLogger()
		API_STATE.LOGGER.Debug("Loaded environment variables from `.env` file successfully: ", &API_STATE)

		API_STATE.InitializeDbHandle()
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

	utils.IsNotFiberChild(func() {
		API_STATE.LOGGER.Infof("Running the API!")
	})

	// Run!
	app.Listen(fmt.Sprintf(":%s", API_STATE.API_PORT))
}
