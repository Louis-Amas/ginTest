package pings

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	pings := r.Group("/pings")
	{
		pings.GET("/", ping)
	}
}
