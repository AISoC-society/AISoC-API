/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/routes/index_route.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2022, Aakash Sen Sharma & Contributors
 */

package routes

import "github.com/gofiber/fiber/v2"

func IndexRoute(c *fiber.Ctx) error {
	return c.JSON("Welcome to AISoC API.")
}
