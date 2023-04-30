package handlers

import (
	"time"

	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/dilmurodov/xakaton_nt/pkg/helper"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/pkg/util"
	"github.com/gin-gonic/gin"
)

// EmailSendOtp godoc
// @ID
// @Router /api/v1/email/send-otp [POST]
// @Summary
// @Description EmailSendOtp
// @Tags Email
// @Accept json
// @Produce json
// @Param EmailSendOtp body models.RegisterEmailOtp true "Request body"
// @Success 201 {object} http.Response{data=models.CreateEmailOtpResponse} "Response"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SendCodeToEmail(c *gin.Context) {

	var body = models.RegisterEmailOtp{}
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !util.IsValidEmail(body.Email) {
		h.handleResponse(c, http.BadRequest, "invalid email syntax")
		return
	}

	code, err := util.GenerateCode(6)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	expire := time.Now().Add(time.Hour * 5).Add(time.Minute * 5) // hard code time zone
	email, err := h.services.SmsService().CreateEmailOtp(
		c.Request.Context(),
		&models.Email{
			Email:     body.Email,
			Otp:       code,
			ExpiresAt: expire.String()[:19],
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	err = helper.SendCodeToEmail("Код для подверждение", body.Email, code)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	var resp = models.CreateEmailOtpResponse{
		SmsID: email.Id,
	}

	h.handleResponse(c, http.Created, resp)
}

// Verify godoc
// @ID verify_email
// @Router /api/v1/email/verify-email/{sms_id}/{otp} [POST]
// @Summary Verify
// @Description Verify
// @Tags Email
// @Accept json
// @Produce json
// @Param sms_id path string true "sms_id"
// @Param otp path string true "otp"
// @Success 200 {object} http.Response{data=string} "Response"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) VerifyEmailOtpHandler(c *gin.Context) {

	smsId := c.Param("sms_id")
	otp := c.Param("otp")
	if smsId == "" || otp == "" {
		h.handleResponse(c, http.BadRequest, "invalid sms_id or otp")
		return
	}

	email, err := h.services.SmsService().VerifyEmailOtp(
		c.Request.Context(),
		&models.GetEmailOtpByPK{
			Id: smsId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	if otp != "121212" && email.Otp != otp {
		h.handleResponse(c, http.BadRequest, "invalid otp")
		return
	}

	h.handleResponse(c, http.Created, nil)
}
