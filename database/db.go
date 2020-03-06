package database

import (
	"fmt"

	"github.com/Louis-Amas/ginTest/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	fmt.Println("Connected to database")

	models.Migrate(db)
	return db, err
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
