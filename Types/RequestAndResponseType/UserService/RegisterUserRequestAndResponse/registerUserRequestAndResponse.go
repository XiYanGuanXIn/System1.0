package RegisterUserRequestAndResponse

import "System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterUserResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}
