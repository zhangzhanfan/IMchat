package router

import (
	"code/service"
	"code/utils"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Logger())
	r.GET("getUserList", service.GetUserList)
	r.POST("createUser", service.CreateUser)
	r.GET("deleteUser", service.DeleteUser)
	r.POST("updateUser", service.UpdateUser)
	r.POST("updateInfo", service.UpdatePhoneAndEmail)
	r.GET("sendMsg", service.SendMsg)
	r.GET("sendUserMsg", service.SendUserMsg)
	// docs.SwaggerInfo.BasePath = ""
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
