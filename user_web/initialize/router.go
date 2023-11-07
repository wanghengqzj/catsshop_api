package initialize

import (
	"github.com/gin-gonic/gin"
	"project/catsshop_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	//TODO

	return Router

}
