/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/utils.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//Run a closure if the current thread is the fiber master thread!
//This function only exists as we use prefork to decrease our operation times.
func IsNotFiberChild(f func()) {
	if !fiber.IsChild() {
		f()
	}
}

//Load the .env file relative to the final executable!
func LoadEnvVars(api_state *APP_STATE) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to read `.env` file.\n%s\n", err)
		os.Exit(1)
	}

	api_state.API_MODE = os.Getenv("API_MODE")
	api_state.API_PORT = os.Getenv("API_PORT")
	api_state.DATABASE_PATH = os.Getenv("DATABASE_PATH")
}
