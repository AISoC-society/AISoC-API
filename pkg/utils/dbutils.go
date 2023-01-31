/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/utils/dbutils.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package utils

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db_singleton          sync.Once
	db_singleton_instance *gorm.DB
)

func GetDbHandle(api_state *APP_STATE) *gorm.DB {
	db_singleton.Do(func() {
		if db, err := gorm.Open(sqlite.Open(api_state.DATABASE_PATH), &gorm.Config{}); err != nil {
			fmt.Printf("Failed to open database file: `%s`.\n%s\n", os.Getenv("DATABASE_PATH"), err.Error())
			os.Exit(1)
		} else {
			db_singleton_instance = db
		}
	})
	return db_singleton_instance
}
