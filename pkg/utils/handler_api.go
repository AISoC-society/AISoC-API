/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/handler_api.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"github.com/aisoc-society/aisoc-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (api_state *APP_STATE) AddUser(ctx *fiber.Ctx) error {
	var user models.LoginCredentials
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Failed to parse request body.", "error": err})
	}

	if user_exists(normal_users, user.Username) || user_exists(admin_users, user.Username) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User already exists!"})
	}

	if bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to salt and hash password", "error": err})
	} else {
		if user.IsAdmin {
			admin_users[user.Username] = string(bytes)
		} else {
			normal_users[user.Username] = string(bytes)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully created the user"})
}

func (api_state *APP_STATE) V1Root(ctx *fiber.Ctx) error {
	return ctx.JSON("Welcome to the AISoC API V1!")
}
