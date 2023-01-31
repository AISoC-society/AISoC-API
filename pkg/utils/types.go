/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/utils.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ENV_VAR_STATE struct {
	API_MODE        string
	API_PORT        string
	DATABASE_PATH   string
	LOGGER          *zap.SugaredLogger
	DATABASE_HANDLE *gorm.DB
}
