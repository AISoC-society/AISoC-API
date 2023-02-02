/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * cmd/api/router/handler/handler.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package handler

import "github.com/gofiber/fiber/v2"

func V1Root(ctx *fiber.Ctx) error {
	return ctx.JSON("Welcome to the AISoC API V1!")
}

func AuthLogin(ctx *fiber.Ctx) error {
	return ctx.JSON("Logged in!")
}
