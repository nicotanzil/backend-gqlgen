package errors

import "errors"

var (
	Register_SentOTP = errors.New("OTP has been sent. Please check your email inbox")
	Auth_Invalid     = errors.New("Invalid Username or Password.")
	JWT_Fail         = errors.New("Fail to create JWT token.")
)