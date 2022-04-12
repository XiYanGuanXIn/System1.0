package User

import "System1.0/Types/ServiceType/UserType"

type User struct {
	UserID      uint              `json:"user_id"`
	Username    string            `json:"user_name"`
	UserType    UserType.UserType `json:"user_type"`
	DisplayName string            `json:"display_name"`
	Password    string            `json:"password"`
	Money       float64           `json:"money"`
}
