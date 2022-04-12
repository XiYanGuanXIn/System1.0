package AuthController

import (
	"net/http"

	"System1.0/Dao/UserDao"
	"System1.0/Service/UserService"
	"System1.0/Types/RequestAndResponseType/Auth/LoginRequestAndResponse"
	"System1.0/Types/RequestAndResponseType/Auth/LogoutRequestAndResponse"
	"System1.0/Types/RequestAndResponseType/Auth/SetUserTypeRequestAndResponse"
	"System1.0/Types/RequestAndResponseType/Auth/WhoAmIRequestAndResponse"
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrorInformation"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (controller AuthController) Login(c *gin.Context) {
	loginRequest := &LoginRequestAndResponse.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// checkUserPassword
	var userService UserService.UserService
	errNo, user, userinfo := userService.CheckPassword(loginRequest.Username, loginRequest.Password)
	loginResponse := &LoginRequestAndResponse.LoginResponse{}
	loginResponse.Code = errNo
	loginResponse.Data.Userinfo = userinfo
	loginResponse.Data.User = user
	loginResponse.Data.Message = ErrorInformation.GenerateErrorInformation(loginResponse.Code)
	c.SetCookie("camp-session", loginRequest.Username, 0, "/", "/", false, false)
	c.JSON(http.StatusOK, loginResponse)
}

func (controller AuthController) Logout(c *gin.Context) {
	logoutResponse := &LogoutRequestAndResponse.LogoutResponse{}

	cookie, err := c.Request.Cookie("camp-session")
	if err == nil {
		c.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		logoutResponse.Code = ErrNo.OK
	} else {
		logoutResponse.Code = ErrNo.LoginRequired
	}
	c.JSON(http.StatusOK, logoutResponse)
}

func (controller AuthController) WhoAmI(c *gin.Context) {
	whoAmIResponse := WhoAmIRequestAndResponse.WhoAmIResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		whoAmIResponse.Code = ErrNo.LoginRequired
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
		c.JSON(http.StatusOK, whoAmIResponse)
		return
	}
	whoAmIResponse.Data.User, err = UserDao.FindUserByUsername(cookie.Value)
	if err != nil {
		whoAmIResponse.Code = ErrNo.UnknownError
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	} else if whoAmIResponse.Data.User.UserID == 0 {
		whoAmIResponse.Code = ErrNo.UserNotExisted
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	} else {
		whoAmIResponse.Code = ErrNo.OK
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	}
	c.JSON(http.StatusOK, whoAmIResponse)
}
func (controller AuthController) SetUserType(c *gin.Context) {
	request := SetUserTypeRequestAndResponse.SetUserTypeRequest{}
	response := SetUserTypeRequestAndResponse.SetUserTypeResponse{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = ErrNo.UnknownError
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	cookie, err := c.Cookie("camp-session");if err!=nil{
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
	}
	response.Code = UserService.UserService{}.SetUserType(cookie,request.UserID,request.UserType)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
