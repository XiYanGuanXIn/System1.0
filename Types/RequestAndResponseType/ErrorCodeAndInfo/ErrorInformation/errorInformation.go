package ErrorInformation

import (
	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
)

func GenerateErrorInformation(code ErrNo.ErrNo) string {
	if code == ErrNo.UnknownError {
		return "An unexpected error has happened.Please try again."
	}
	if code == ErrNo.WrongPassword {
		return "Wrong Password!"
	}
	if code == ErrNo.LoginRequired {
		return "Your information is out of date.Please log in first."
	}
	if code == ErrNo.UserNotExisted {
		return "There is no such user."
	}
	if code == ErrNo.UserHasExisted {
		return "This username has been registered before."
	}
	if code == ErrNo.PermDenied {
		return "Sorry. You don't have permission to do so."
	}
	if code == ErrNo.MoneyNotEnough {
		return "Your account doesn't have enough money to pay this bill!"
	}
	if code == ErrNo.ParamInvalid {
		return "The parameters are invalid!"
	}
	return "Success!"
}
