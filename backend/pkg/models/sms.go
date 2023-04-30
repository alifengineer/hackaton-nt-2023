package models

type Sms struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Text        string `json:"text"`
	Recipient   string `json:"recipient"`
	ExpiresAt   string `json:"expires_at"`
	Otp         string `json:"otp"`
	PhoneNumber string `json:"phone_number"`
	SendCount   int64  `json:"send_count"`
}

type ConfirmOtpRequest struct {
	SmsID string `json:"sms_id"`
	Otp   string `json:"otp"`
}

type SendOtpRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type SendOtpResponse struct {
	SmsID string `json:"sms_id"`
}

type Content struct {
	Text string `json:"text"`
}

type SMS struct {
	Originator string  `json:"originator"`
	Content    Content `json:"content"`
}

type Message struct {
	Recipient string `json:"recipient"`
	MessageID string `json:"message-id"`
	SMS       SMS    `json:"sms"`
}

type Body struct {
	Messages []Message `json:"messages"`
}
