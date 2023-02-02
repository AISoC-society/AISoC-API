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

//Print log information and call os.Exit with an exit code of 1.
func (api_state *APP_STATE) panic(template string, args ...interface{}) {
	api_state.LOGGER.Errorf(template, args)
	os.Exit(1)
}

//Load the `.env` file relative to the final executable!
func (api_state *APP_STATE) LoadEnvVars() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to read `.env` file.\n%s\n", err)
		os.Exit(1)
	}

	api_state.API_MODE = os.Getenv("API_MODE")
	api_state.API_PORT = os.Getenv("API_PORT")
	api_state.DATABASE_PATH = os.Getenv("DATABASE_PATH")
}
