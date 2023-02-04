/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/models/models.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
	Credits         uint
}

//JWT generation requirements.
type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type JwtClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}
