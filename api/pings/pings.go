import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	pings := r.Group("/posts")
	{
		pings.GET("/", list)
	}
}