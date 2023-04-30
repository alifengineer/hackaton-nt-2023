package handlers

import (
	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLeagues(c *gin.Context) {

	resp, err := h.services.MatchService().GetAllLeagues(
		c.Request.Context(),
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetMatchesByTUR(c *gin.Context) {

	var body models.GetMatchesByTURStatusRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.MatchService().GetMatchesByTURStatus(
		c.Request.Context(),
		&body,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetTopTeamsInLeague(c *gin.Context) {

	var body models.GetTopTeamsInLeagueRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.MatchService().GetTopTeamsInLeague(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllMatches(c *gin.Context) {

	var body models.GetAllMatchRequest

	leagueID := c.Query("league_id")
	if leagueID != "" {
		body.LeagueID = leagueID
	}
	seasonID := c.Query("season_id")
	if seasonID != "" {
		body.SeasonID = seasonID
	}
	mDateFrom := c.Query("m_date_from")
	if mDateFrom != "" {
		body.MDateFrom = mDateFrom
	}
	mDateTo := c.Query("m_date_to")
	if mDateTo != "" {
		body.MDateTo = mDateTo
	}
	status := c.Query("status")
	if status != "" {
		body.Status = status
	}
	teamID := c.Query("team_id")
	if teamID != "" {
		body.TeamID = teamID
	}
	search := c.Query("search")
	if search != "" {
		body.Search = search
	}

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

	resp, err := h.services.MatchService().GetAllMatch(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) CreateMatch(c *gin.Context) {

	var body models.Match
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.MatchService().CreateMatch(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetMatchByID(c *gin.Context) {

	var body models.GetMatchByIDRequest

	matchID := c.Param("id")
	if matchID == "" {
		h.handleResponse(c, http.BadRequest, "match_id is required")
		return
	}

	body.ID = matchID
	resp, err := h.services.MatchService().GetMatchByID(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetLeagueByID(c *gin.Context) {

	var body models.GetLeagueByIDRequest

	leagueID := c.Param("id")
	if leagueID == "" {
		h.handleResponse(c, http.BadRequest, "league_id is required")
		return
	}

	body.ID = leagueID
	resp, err := h.services.MatchService().GetLeagueByID(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllTeams(c *gin.Context) {

	var body models.GetAllTeamsRequest

	leagueID := c.Query("league_id")
	if leagueID != "" {
		body.LeagueID = leagueID
	}
	seasonID := c.Query("season_id")
	if seasonID != "" {
		body.SeasonID = seasonID
	}
	search := c.Query("search")
	if search != "" {
		body.Search = search
	}

	resp, err := h.services.MatchService().GetAllTeams(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetTeamByID(c *gin.Context) {

	var body models.GetTeamByIDRequest

	teamID := c.Param("id")
	if teamID == "" {
		h.handleResponse(c, http.BadRequest, "team_id is required")
		return
	}

	body.ID = teamID
	resp, err := h.services.MatchService().GetTeamByID(
		c.Request.Context(),
		&body,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetAllSeasons(c *gin.Context) {

	var body models.GetAllSeasonsRequest

	leagueID := c.Query("id")
	if leagueID != "" {
		body.LeagueID = leagueID
	}

	resp, err := h.services.MatchService().GetAllSeasons(
		c.Request.Context(),
		&body,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
