package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/jwt"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	var result models.HasAccessModel
	if ok := h.hasAccess(c, &result); !ok {
		c.Abort()
		return
	}

	c.Set("auth", &result)

	c.Next()
}

func (h *Handler) hasAccess(c *gin.Context, result *models.HasAccessModel) bool {

	bearerToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 || strArr[0] != "Bearer" {
		h.handleResponse(c, http.Forbidden, "token error: wrong format")
		return false
	}
	accessToken := strArr[1]

	claims, err := jwt.ExtractClaims(accessToken, config.SigningKey)
	if err != nil {
		h.handleResponse(c, http.Forbidden, "no access")
		return false
	}

	var tm time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}

	if tm.Unix() < time.Now().Unix() {
		h.handleResponse(c, http.Forbidden, "token expired")
		return false
	}

	userId := claims["user_id"]
	if userId == nil {
		h.handleResponse(c, http.Forbidden, "no access")
		return false
	}

	email := claims["email"]
	if email == nil {
		h.handleResponse(c, http.Forbidden, "no access")
		return false
	}

	_, err = h.services.UserService().GetUserByID(c, &models.GetUserByIDRequest{
		UserId: userId.(string),
		Email:  email.(string),
	})
	if err != nil {
		h.handleResponse(c, http.Forbidden, err.Error())
		return false
	}

	result.UserId = userId.(string)
	result.Email = email.(string)

	return true
}
