package LoginRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/ServiceType/User"
)

type LoginRequest struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
type LoginResponse struct {
	Code ErrNo.ErrNo `json:"Code" binding:"required"`
	Data struct {
		Message  string `json:"message" binding:"required"`
		User     User.User
		Userinfo string `json:"userinfo" `
	}
}
