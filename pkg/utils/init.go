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
	"strings"

	"go.uber.org/zap"
)

//Initialize the uber-go-zap logger according to the `API_MODE` environment variable (production/debug).
func InitializeLogger(api_state *ENV_VAR_STATE) {
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
	return
}

//Initialize the database handle in the api_state
func InitializeDbHandle(api_state *ENV_VAR_STATE) {
	api_state.DATABASE_HANDLE = GetDbHandle(api_state)
}
