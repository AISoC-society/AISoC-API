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
	API_MODE        string
	API_PORT        string
	DATABASE_PATH   string
	LOGGER          *zap.SugaredLogger
	DATABASE_HANDLE *gorm.DB
}

//Schema models.
type Event struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	name      string
}

type Member struct {
	ID              uint `gorm:"primaryKey;autoIncrement:1"`
	Email           string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime:milli"`
	EventsAttended  []Event   `gorm:"foreignkey:ID"`
	AttendanceCount int
}
