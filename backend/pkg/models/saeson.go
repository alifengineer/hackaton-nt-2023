package models

type Season struct {
	ID         string
	SeasonFrom int
	SeasonTo   int
	LeagueID   string
	IsCurrent  bool
}

type GetAllSeasonsRequest struct {
	LeagueID string
}