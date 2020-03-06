package main

import (
	"github.com/Louis-Amas/ginTest/api"
	"github.com/Louis-Amas/ginTest/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	r.Use(database.Inject(db))
	api.ApplyRoutes(r)
	r.Run(":8080")
}
