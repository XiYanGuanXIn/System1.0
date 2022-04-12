package FindUserInfoRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/ServiceType/User"
)

type FindUserInfoRequest struct {
}

type FindUserInfoResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		User     User.User
		UserInfo string `json:"userInfo"`
		Message  string `json:"message"`
	}
}
