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
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var API_STATE utils.APP_STATE

func main() {
	// Since we prefork for performance boost, we need to keep master thread checks!
	utils.IsNotFiberChild(func() {
		utils.LoadEnvVars(&API_STATE)
		utils.InitializeLogger(&API_STATE)

		API_STATE.LOGGER.Infoln("Loaded environment variables from `.env` file successfully!")
		bytes, _ := json.Marshal(API_STATE)
		API_STATE.LOGGER.Infoln(string(bytes))

		utils.InitializeDbHandle(&API_STATE)
	})

	// Creating the API.
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Go-Fiber",
		AppName:       "AISoC API",
	})

	// Registering all routes.
	router.SetupRoutes(app)
	utils.IsNotFiberChild(func() { API_STATE.LOGGER.Infoln("Successfully setup all API Routes!") })

	// Run!
	app.Listen(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
