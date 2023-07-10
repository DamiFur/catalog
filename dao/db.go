package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *sql.DB
	envs = map[string]string{"prod": "Catalog", "test": "Test"}
)

const (
	dsn = "root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func CreateDatabase(env string) {
	g, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn, envs[env])), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to database: %v", err))
	}

	g.AutoMigrate(&Category{}, &Item{})

	dbSql, err := g.DB()
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to database: %v", err))
	}

	dbSql.SetMaxIdleConns(5)
	dbSql.SetMaxOpenConns(100)
	dbSql.SetConnMaxLifetime(time.Minute * 30)

	db = dbSql
}

func CloseDB() {
	db.Close()
}

func GetSession() (*gorm.DB, error) {
	grm, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)

	return grm.WithContext(ctx), nil

}
