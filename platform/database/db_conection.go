package database

import (
	"dont/hexagonal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// connectDb
func ConnectDb() *gorm.DB {
	// Define database connection settings.
	db_type := os.Getenv("DB_TYPE")
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder(db_type)
	if err != nil {
		panic(err)
	}
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/

	db, err := gorm.Open(mysql.Open(mysqlConnURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",                          // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                        // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                       // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("", ""), // use name replacer to change struct/field name before convert it to db name
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	dbConfig, _ := db.DB()
	dbConfig.SetMaxOpenConns(maxConn)
	dbConfig.SetMaxIdleConns(maxIdleConn)
	dbConfig.SetConnMaxLifetime(time.Duration(maxLifetimeConn))
	return db
}
