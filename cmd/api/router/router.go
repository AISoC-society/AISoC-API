/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/router/router.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package router

import (
	"api/cmd/api/router/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")
	v1.Get("/", handler.V1_root)
}
