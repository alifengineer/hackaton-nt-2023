package handlers

import (
	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateNews(c *gin.Context) {

	var body models.CreateNewsRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.NewsService().CreateNews(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllNews(c *gin.Context) {

	var body models.GetAllNewRequest

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	body.Limit = limit

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	body.Offset = offset

	resp, err := h.services.NewsService().GetAllNews(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetNewsByID(c *gin.Context) {

	var body models.GetNewsByIDRequest

	newsID := c.Param("id")
	body.NewsID = newsID

	resp, err := h.services.NewsService().GetNewsByID(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
