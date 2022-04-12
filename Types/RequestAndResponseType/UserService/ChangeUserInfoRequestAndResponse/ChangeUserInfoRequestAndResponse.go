package ChangeUserInfoRequestAndResponse

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
)

type ChangeUserInfoRequest struct {
	NewPassword string `json:"new_password"`
	DisplayName string `json:"display_name"`
	UserInfo    string `json:"user_info"`
}

type ChangeUserInfoResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
