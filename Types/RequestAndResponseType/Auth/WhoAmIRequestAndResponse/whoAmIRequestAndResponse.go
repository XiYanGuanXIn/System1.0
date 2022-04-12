package WhoAmIRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/ServiceType/User"
)

type WhoAmIRequest struct {
}

type WhoAmIResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		User    User.User
		Message string `json:"message"`
	}
}
