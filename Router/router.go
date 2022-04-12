package Router

import (
	"System1.0/Controller/AuthController"
	"System1.0/Controller/UserServiceController"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	// auth
	g.POST("/auth/login", AuthController.AuthController{}.Login)
	g.POST("/auth/logout", AuthController.AuthController{}.Logout)
	g.POST("/auth/whoAmI", AuthController.AuthController{}.WhoAmI)
	g.POST("/auth/setUserType", AuthController.AuthController{}.SetUserType)
	//user
	g.POST("/user/changeUserInfo", UserServiceController.UserServiceController{}.ChangeUserInfo)
	g.POST("/user/register", UserServiceController.UserServiceController{}.RegisterUser)
	g.POST("/user/addMoney", UserServiceController.UserServiceController{}.AddMoney)
	g.POST("/user/getUserInfo", UserServiceController.UserServiceController{}.FindUserInfo)
}
