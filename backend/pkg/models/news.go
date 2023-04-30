package models

type News struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	IsImportant bool   `json:"is_important"`
}

type GetNewsByIDRequest struct {
	NewsID string `json:"news_id"`
}

type GetNewsByIDResponse struct {
	News *News `json:"news"`
}

type GetAllNewRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetAllNewsResponse struct {
	NewsList   []*News `json:"news_list"`
	TotalCount int     `json:"total_count"`
}

type CreateNewsRequest struct {
	Title   string `json:"title"`
	Image   string `json:"image"`
	Content string `json:"content"`
}