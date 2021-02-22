package errorlist

import "errors"

var (
	Register_SentOTP 	= errors.New("OTP has been sent. Please check your email inbox")
	Auth_Invalid     	= errors.New("Invalid Username or Password.")
	JWT_Fail         	= errors.New("Fail to create JWT token.")
	User_NonExists 		= errors.New("User not found.")
	Access_Denied		= errors.New("Access Denied.")
	Account_Suspended 	= errors.New("Account Suspended.")
)