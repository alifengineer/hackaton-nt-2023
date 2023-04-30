package models

type League struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	PrevTur int    `json:"prev_tur"`
	NextTur int    `json:"next_tur"`
}

type Tur struct {
	ID         string `json:"id"`
	LeagueID   string `json:"league_id"`
	Soni       int    `json:"soni"`
	CurrentTur int    `json:"current_tur"`
}

type GetTopTeamsInLeagueRequest struct {
	LeagueID string `json:"league_id"`
}

type GetTopTeamsInLeagueResponse struct {
	Teams []*Team `json:"teams"`
}

type GetTURByLeagueIDRequest struct {
	LeagueID string `json:"league_id"`
}

type GetTURByLeagueIDResponse struct {
	TUR *Tur `json:"tur"`
}

type GetLeagueByIDRequest struct {
	ID string `json:"id"`
}

type GetLeagueByIDResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Image  string  `json:"image"`
	Tur    *Tur    `json:"tur"`
	Season *Season `json:"season"`
}

type GetLeagueSeasonTeamsRequest struct {
	LeagueID string `json:"league_id"`
	SeasonID string `json:"season_id"`
}

type GetLeagueSeasonTeamsResponse struct {
	Teams []*Team `json:"teams"`
}

type CreateLeagueSeasonTeamsRequest struct {
	LeagueID string  `json:"league_id"`
	SeasonID string  `json:"season_id"`
	Teams    []*Team `json:"teams"`
}
