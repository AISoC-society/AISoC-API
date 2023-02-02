/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/router/router.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package router

import (
	"github.com/aisoc-society/aisoc-api/cmd/api/router/handler"
	"github.com/aisoc-society/aisoc-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, api_state *utils.APP_STATE) {
	// General API endpoints.
	api := app.Group("/api", logger.New())

	// Authentication endpoints.
	auth := api.Group("/auth")
	auth.Post("/login", handler.AuthLogin)

	// Version 1 endpoints.
	// NOTE: Over further releases we can retire old versions and bump them
	// semantically here to maintain backwards compatibility.
	v1 := api.Group("/v1")
	v1.Get("/", handler.V1Root)

	utils.IsNotFiberChild(func() {
		api_state.LOGGER.Infoln("Successfully setup all API route handlers!")
	})
}
