package LogoutRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
)

type LogoutRequest struct {
}
type LogoutResponse struct {
	Code ErrNo.ErrNo `json:"Code" required:"True"`
}
