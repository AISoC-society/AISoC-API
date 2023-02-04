/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/types.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type APP_STATE struct {
	API_MODE                   string
	API_PORT                   string
	DATABASE_PATH              string
	MAX_CONCURRENT_CONNECTIONS int
	EXPIRATION_TIME            time.Duration
	LOGGER                     *zap.SugaredLogger
	DATABASE_HANDLE            *gorm.DB
}
