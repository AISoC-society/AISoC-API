/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/handler_auth.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"os"

	"time"

	"github.com/aisoc-society/aisoc-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var normal_users = map[string]string{}
var admin_users = map[string]string{}

func (api_state *APP_STATE) AuthLogin(ctx *fiber.Ctx) error {
	var credentials models.LoginCredentials

	if err := ctx.BodyParser(&credentials); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Failed to parse request body.", "error": err})
	}

	// We need to override whatever the user set as `is_admin` in the request body.
	// This field is kept around as it is used to unmarshal the adduser route request body as well.
	credentials.IsAdmin = false
	if !is_valid_user(normal_users, &credentials) {
		if !is_valid_user(admin_users, &credentials) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not registered!"})
		} else {
			credentials.IsAdmin = true
		}
	}

	issued_at := time.Now()
	expires_at := issued_at.Add(time.Minute * 5)

	claims := &models.JwtClaims{
		Username: credentials.Username,
		IsAdmin:  credentials.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Audience:  "API-Consumer-Client",
			ExpiresAt: expires_at.Unix(),
			IssuedAt:  issued_at.Unix(),
			Issuer:    "AISoC-API",
			NotBefore: issued_at.Unix(),
			Subject:   "AISoC-API JWT-auth-key",
		},
	}

	if token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))); err != nil {
		if err := ctx.SendStatus(fiber.StatusInternalServerError); err != nil {
			return err
		}
		return err
	} else {
		return ctx.JSON(fiber.Map{"token": token})
	}
}

//Check if a user exists in our records and accordingly compare our stored credentials with the provided ones.
func is_valid_user(hash_map map[string]string, credentials *models.LoginCredentials) (valid bool) {
	if user_exists(hash_map, credentials.Username) {
		stored_pass := hash_map[credentials.Username]
		if bcrypt.CompareHashAndPassword([]byte(stored_pass), []byte(credentials.Password)) == nil {
			valid = true
		}
	}
	return
}

//Check if a username exists in a hash_map.
func user_exists(hash_map map[string]string, username string) (exists bool) {
	_, exists = hash_map[username]
	return
}
