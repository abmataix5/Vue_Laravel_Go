package Routes

import (
	"first-api/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	grp1 := r.Group("/api")
	{
		grp1.GET("user/", Controllers.GetWorkers)
		grp1.POST("user/", Controllers.CreateWorker)
		grp1.GET("user/:id", Controllers.GetWorkerByID)
		grp1.PUT("user/:id", Controllers.UpdateWorker)
		grp1.DELETE("user/:id", Controllers.DeleteWorker)
	}
	return r

}
