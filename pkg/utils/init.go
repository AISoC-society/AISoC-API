/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/init.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

//Initialize the global logger.
func (api_state *APP_STATE) InitializeLogger() {
	var zap_logger *zap.Logger
	var err error

	if strings.Compare(strings.ToLower(api_state.API_MODE), "production") == 0 {
		zap_logger, err = zap.NewProduction()
	} else {
		zap_logger, err = zap.NewDevelopment()
	}

	if err != nil {
		fmt.Printf("Failed to initialize logger: %s\n", err)
		os.Exit(1)
	}

	defer zap_logger.Sync()
	api_state.LOGGER = zap_logger.Sugar()
	api_state.LOGGER.Debug("Successfully initialized zap-logger!")
	return
}

//Initialize the global database handle.
func (api_state *APP_STATE) InitializeDbHandle() {
	api_state.DATABASE_HANDLE = GetDbHandle(api_state)
	api_state.LOGGER.Debug("Successfully established database handshake!")

	if err := api_state.DATABASE_HANDLE.AutoMigrate(&Member{}, &Event{}); err != nil {
		api_state.panic("Schema auto migration failed!")
	}
	api_state.LOGGER.Debug("Successfully completed schema migration!")
}
