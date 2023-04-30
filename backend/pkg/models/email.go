package models

type Email struct {
	Email     string `json:"email"`
	Otp       string `json:"otp"`
	ExpiresAt string `json:"expires_at"`
	Id        string `json:"id"`
}

type GetEmailOtpByPK struct {
	Id string `json:"id"`
}

type RegisterEmailOtp struct {
	Email string `json:"email"`
}

type Verify struct {
	SmsId string `json:"sms_id"`
	Otp   string `json:"otp"`
}

type CreateEmailOtpResponse struct {
	SmsID string `json:"sms_id"`
}
