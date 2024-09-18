package request

import "github.com/gookit/validate"

type LoginForm struct {
	EmailOrPhone string `json:"emailOrPhone" xml:"emailOrPhone" form:"emailOrPhone" validate:"required"`
	PIN          string `json:"pin" xml:"pin" form:"pin" validate:"required|len:6"`
}

// Messages you can custom validator error messages.
func (f LoginForm) Messages() map[string]string {
	return validate.MS{
		"EmailOrPhone.required": "Email atau nomor telepon harus diisi.",
		"PIN.required":          "PIN harus diisi.",
		"PIN.len":               "PIN harus tepat 6 karakter.",
	}
}

// Translates you can custom field translates.
func (f LoginForm) Translates() map[string]string {
	return validate.MS{
		"EmailOrPhone": "Email atau Nomor Telepon",
		"PIN":          "PIN",
	}
}

type CheckUserForm struct {
	EmailOrPhone string `json:"emailOrPhone" xml:"emailOrPhone" form:"emailOrPhone" validate:"required"`
}

// Messages you can custom validator error messages.
func (f CheckUserForm) Messages() map[string]string {
	return validate.MS{
		"EmailOrPhone.required": "Email atau nomor telepon harus diisi.",
	}
}

// Translates you can custom field translates.
func (f CheckUserForm) Translates() map[string]string {
	return validate.MS{
		"EmailOrPhone": "Email atau Nomor Telepon",
	}
}

type CheckOTPForm struct {
	EmailOrPhone string `json:"emailOrPhone" xml:"emailOrPhone" form:"emailOrPhone" validate:"required"`
	OTP          string `json:"otp" xml:"otp" form:"otp" validate:"required"`
}

// Messages you can custom validator error messages.
func (f CheckOTPForm) Messages() map[string]string {
	return validate.MS{
		"EmailOrPhone.required": "Email atau nomor telepon harus diisi.",
		"OTP.required":          "OTP harus diisi.",
	}
}

// Translates you can custom field translates.
func (f CheckOTPForm) Translates() map[string]string {
	return validate.MS{
		"EmailOrPhone": "Email atau Nomor Telepon",
		"OTP":          "OTP",
	}
}

type ResendOTPForm struct {
	EmailOrPhone string `json:"emailOrPhone" xml:"emailOrPhone" form:"emailOrPhone" validate:"required"`
}

// Messages you can custom validator error messages.
func (f ResendOTPForm) Messages() map[string]string {
	return validate.MS{
		"EmailOrPhone.required": "Email atau nomor telepon harus diisi.",
	}
}

// Translates you can custom field translates.
func (f ResendOTPForm) Translates() map[string]string {
	return validate.MS{
		"EmailOrPhone": "Email atau Nomor Telepon",
	}
}
