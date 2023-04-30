package models

type Match struct {
	ID         string `json:"id"`
	HomeTeamID string `json:"home_team_id"`
	AwayTeamID string `json:"away_team_id"`
	LeagueID   string `json:"league_id"`
	SeasonID   string `json:"season_id"`
	HomeScore  int    `json:"home_score"`
	AwayScore  int    `json:"away_score"`
	MDate      string `json:"m_date"`
	Status     string `json:"status"`
	Tur        string `json:"tur"`
}

type GetMatchesByTURStatusRequest struct {
	Tur      int    `json:"tur"`
	Status   string `json:"status"`
	LeagueID string `json:"league_id"`
}

type MatchTeam struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Score int    `json:"score"`
}

type TeamsWithResults struct {
	ID         string     `json:"id"`
	HomeTeam   *MatchTeam `json:"home_team"`
	AwayTeam   *MatchTeam `json:"away_team"`
	MDate      string     `json:"m_date"`
	Winner     string     `json:"winner"`
	Status     string     `json:"status"`
	Tur        int        `json:"tur"`
	LeagueID   string     `json:"league_id"`
	LeagueName string     `json:"league_name"`
	SeasonID   string     `json:"season_id"`
	SeasonFrom int        `json:"season_year_from"`
	SeasonTo   int        `json:"season_year_to"`
}

type GetMatchesByTURStatusResponse struct {
	Matches []*TeamsWithResults `json:"matches"`
}

type GetAllMatchRequest struct {
	LeagueID  string `json:"league_id"`
	SeasonID  string `json:"season_id"`
	Tur       int    `json:"tur"`
	MDateFrom string `json:"m_date_from"`
	MDateTo   string `json:"m_date_to"`
	Status    string `json:"status"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	TeamID    string `json:"team_id"`
	Search    string `json:"search"`
}

type GetAllMatchResponse struct {
	Matches    []*TeamsWithResults `json:"matches"`
	TotalCount int                 `json:"total_count"`
}

type GetMatchByIDRequest struct {
	ID string `json:"id"`
}

type GetMatchByIDResponse struct {
	Match *TeamsWithResults `json:"match"`
}

// // topnew api
// const news = [
// 	{
// 		"id": "uuid",
// 		"title": "title",
// 		"image": "https://upload.wikimedia.org/wikipedia/ru/thumb/4/47/FC_Barcelona_%28crest%29.svg/1200px-FC_Barcelona_%28crest%29.svg.png",
// 		"date": "2020-10-10",
// 	},
// ]

// // get all news

// const news = [
// 	{
// 		"id": "uuid",
// 		"title": "title",
// 		"image": "https://upload.wikimedia.org/wikipedia/ru/thumb/4/47/FC_Barcelona_%28crest%29.svg/1200px-FC_Barcelona_%28crest%29.svg.png",
// 		"content": "content",
// 		"date": "2020-10-10",
// 	},
// ]
