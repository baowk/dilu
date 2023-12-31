package codes

var enUSText = map[int]string{
	SUCCESS:                "OK",
	FAILURE:                "FAIL",
	NotFound_404:           "resources not found",
	RequestsFrequently:     "Too many requests",
	InvalidParameter:       "Parameter error",
	AuthorizationError_403: "You have no permission",
	InvalidToken_401:       "Invalid Token",
	CaptchaErr:             "Get captcha fail,Please try again later",
	CaptchaVerifyErr:       "Verification code error",
	PhoneExistErr:          "Phone already exists",
	EmailExistErr:          "Email already exists",
	UserNotExist:           "Current account does not exist, please register first",
	UserLock:               "Current account is frozen",
	PwdNotExist:            "Password not yet set, please select the verification code to log in",
	ThirdNotScan:           "No scanning information received",
	ThirdExpire:            "expired",
	ErrRePassword:          "Repeat Password Inconsistency",
	ErrPasswordFMT:         "Wrong password format",
	ErrMobileOrEmail:       "Must have a cell phone number or email to register",
	ErrVerifyCode:          "CAPTCHA error",
	ErrBind:                "Binding Failure",
	ErrUserExist:           "Account has been registered, please login directly",
	ErrUsernameOrPwd:       "Wrong account or password",
	ErrPwd:                 "Wrong  password",
}
