package SetUserTypeRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/ServiceType/UserType"
)

type SetUserTypeRequest struct {
	UserID   uint              `json:"user_id"`
	UserType UserType.UserType `json:"user_type"`
}
type SetUserTypeResponse struct {
	Code ErrNo.ErrNo `json:"code"`
	Data struct {
		Message string `json:"message"`
	}
}
