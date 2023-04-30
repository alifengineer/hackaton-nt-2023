package handlers

import (
	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/customerrors"
	"github.com/dilmurodov/xakaton_nt/pkg/jwt"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/pkg/security"
	"github.com/dilmurodov/xakaton_nt/pkg/util"
	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @ID user_register
// @Router /api/v1/auth/register [POST]
// @Summary Register User
// @Description Register User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.RegisterUserRequest true "Request body"
// @Success 201 {object} http.Response{data=models.UserWithAuth} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RegisterHandler(c *gin.Context) {
	var user *models.RegisterUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if len(user.Password) < 6 {
		h.handleResponse(c, http.BadRequest, "password must be at least 6 characters")
		return
	}

	if !util.IsValidEmail(user.Email) {
		h.handleResponse(c, http.BadRequest, "invalid email")
		return
	}

	_, err = h.services.UserService().GetUserPasswordByEmail(c.Request.Context(), user.Email)
	if err != nil {
		if e, ok := err.(*customerrors.UserNotFoundWithEmailError); !ok {
			h.handleResponse(c, http.InternalServerError, e.Error())
			return
		}
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUser(
		c.Request.Context(),
		&models.CreateUserRequest{
			User: &models.User{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Password:  hashedPassword,
			},
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	m := map[interface{}]interface{}{
		"user_id": resp.Guid,
		"email":   resp.Email,
	}

	accessToken, refreshTokenk, err := jwt.GenJWT(m, []byte(config.SigningKey))
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, &models.UserWithAuth{
		User: &models.User{
			Guid:      resp.Guid,
			Email:     resp.Email,
			FirstName: resp.FirstName,
			LastName:  resp.LastName,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshTokenk,
	})
}

// LoginUser godoc
// @ID login_user
// @Router /api/v1/auth/login [POST]
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.LoginUserRequest true "Request body"
// @Success 201 {object} http.Response{data=models.UserWithAuth} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) LoginHandler(c *gin.Context) {

	var req *models.LoginUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().GetUserByCredentials(
		c.Request.Context(),
		&models.GetByCredentialsRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	m := map[interface{}]interface{}{
		"user_id": resp.Guid,
		"email":   resp.Email,
	}
	accessToken, refreshTokenk, err := jwt.GenJWT(m, []byte(config.SigningKey))
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, &models.UserWithAuth{
		User: &models.User{
			Guid:      resp.Guid,
			Email:     resp.Email,
			FirstName: resp.FirstName,
			LastName:  resp.LastName,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshTokenk,
	})
}
