package errors

import "errors"

var Error_RegisterOTPSent = errors.New("OTP has been sent. Please check your email inbox")
var Auth_Invalid = errors.New("Invalid Username/Password ")