/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/router/router.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package router

import (
	"github.com/aisoc-society/aisoc-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App, api_state *utils.APP_STATE) {
	//Live resources monitor.
	if api_state.API_MODE == "debug" {
		app.Get("/monitor", monitor.New(monitor.Config{Title: "AISoC API Live Monitor"}))
	}

	// General API endpoints.
	api := app.Group("/api", logger.New())

	// Authentication endpoints.
	{
		auth := api.Group("/auth")
		auth.Get("/login", api_state.AuthLogin)
	}

	// Version 1 endpoints.
	{
		v1 := api.Group("/v1")
		v1.Get("/", api_state.V1Root)
		v1.Put("/adduser", api_state.AddUser)
	}

	utils.IsNotFiberChild(func() {
		api_state.LOGGER.Infoln("Successfully setup all API route handlers!")
	})
}
