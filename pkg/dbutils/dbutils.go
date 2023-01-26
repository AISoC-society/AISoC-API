/* SPDX-License-Identifier: BSD 2-Clause "Simplified" License
 *
 * pkg/dbutils/dbutils.go
 *
 * Created by:	Aakash Sen Sharma, January 2023
 * Copyright:	(C) 2023, Aakash Sen Sharma & Contributors
 */

package dbutils

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once         sync.Once
	db_singleton *gorm.DB
)

func GetDbHandle() *gorm.DB {
	once.Do(func() {
		if db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_PATH")), &gorm.Config{}); err != nil {
			fmt.Printf("Failed to open database file: `%s`.\n%s\n", os.Getenv("DATABASE_PATH"), err.Error())
			os.Exit(1)
		} else {
			db_singleton = db
		}
	})
	return db_singleton
}
