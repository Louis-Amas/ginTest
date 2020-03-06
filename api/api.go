package api

import (
	"github.com/Louis-Amas/ginTest/api/pings"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		pings.ApplyRoutes(api)
	}
}
