package route

import (
	v1 "example_server/apis/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter use to initialize router
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	apisV1 := router.Group("/api/v1")
	{
		// CRUD on operations api
		apisV1.GET("/testapi", v1.Testapi)
		//users.POST("/jsondata", v1.jsondata)
	}
	return router
}