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
	"os"

	"github.com/aisoc-society/aisoc-api/cmd/api/router"
	"github.com/aisoc-society/aisoc-api/pkg/utils"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	// NOTE: Prefork seems to be finnicky but fast! Prefork guards are already setup.
	// NOTE: Sonic is used for marshaling and unmarshaling as it is an optimized simd json codec.
	app := fiber.New(fiber.Config{
		Prefork:       false,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Go-Fiber",
		AppName:       "AISoC API",
	})

	// Ratelimit configuration.
	app.Use(limiter.New(limiter.Config{
		Max:        API_STATE.MAX_CONCURRENT_CONNECTIONS,
		Expiration: API_STATE.EXPIRATION_TIME,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
	}))

	// Ignore favicon requests.
	app.Use(favicon.New())

	// Recover middleware.
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	// Registering all routes.
	router.SetupRoutes(app, &API_STATE)
	utils.IsNotFiberChild(func() {
		api_stack, _ := sonic.Marshal(app.Stack())
		API_STATE.LOGGER.Debugf("API-Routes: %s", string(api_stack))

		API_STATE.LOGGER.Infof("Running the API!")
	})

	// Run!
	if err := app.Listen(fmt.Sprintf(":%s", API_STATE.API_PORT)); err != nil {
		fmt.Printf("Failed to start Go-Fiber API: %s\n", err)
		os.Exit(1)
	}
}
