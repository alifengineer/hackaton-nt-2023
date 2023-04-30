package models

type Team struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Score      int    `json:"score"`
	Scored     int    `json:"scored"`
	Missed     int    `json:"missed"`
	Difference int    `json:"difference"`
	Played     int    `json:"played"`
}

type GetAllTeamsRequest struct {
	LeagueID string `json:"league_id"`
	SeasonID string `json:"season_id"`
	Search   string `json:"search"`
}

type GetAllTeamsResponse struct {
	Teams      []*Team `json:"teams"`
	TotalCount int     `json:"total_count"`
}

type GetTeamByIDRequest struct {
	ID string `json:"id"`
}

type GetTeamByIDResponse struct {
	Team *Team `json:"team"`
}
