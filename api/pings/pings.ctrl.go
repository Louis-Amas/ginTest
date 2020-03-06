package pings

import (
	"github.com/Louis-Amas/ginTest/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ping(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	u := models.User{}
	db.Last(&u)

	c.JSON(200, u)
}
