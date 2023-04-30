package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
	"github.com/google/uuid"
)

type newsRepo struct {
	db *sql.DB
}

func NewNewsRepo(db *sql.DB) storage.NewsRepoI {
	return &newsRepo{
		db: db,
	}
}

func (n *newsRepo) CreateNews(ctx context.Context, req *models.CreateNewsRequest) (resp *models.News, err error) {

	id := uuid.NewString()

	query := `
		INSERT INTO news(
			id,
			title,
			content,
			image
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING id, title, content, image, created_at`
	err = n.db.QueryRowContext(ctx, query,
		id,
		req.Title,
		req.Content,
		req.Image).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Content,
		&resp.Image,
		&resp.CreatedAt,
	)
	if err != nil {

		return nil, err
	}
	return resp, err
}

func (n *newsRepo) GetAllNews(ctx context.Context, req *models.GetAllNewRequest) (resp *models.GetAllNewsResponse, err error) {

	resp = &models.GetAllNewsResponse{
		TotalCount: 0,
		NewsList:   []*models.News{},
	}
	query := `
		SELECT
			id,
			title,
			content,
			image,
			created_at,
			COUNT(1) FILTER (WHERE deleted_at IS NULL) OVER () AS total_count
		FROM news
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	rows, err := n.db.QueryContext(ctx, query,
		req.Limit,
		req.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r models.News
		err = rows.Scan(
			&r.ID,
			&r.Title,
			&r.Content,
			&r.Image,
			&r.CreatedAt,
			&resp.TotalCount,
		)
		if err != nil {
			return nil, err
		}
		resp.NewsList = append(resp.NewsList, &r)
	}
	return resp, err
}

func (n *newsRepo) GetLatesNews(ctx context.Context, req *models.GetAllNewRequest) (resp *models.GetAllNewsResponse, err error) {

	resp = &models.GetAllNewsResponse{
		TotalCount: 0,
		NewsList:   []*models.News{},
	}
	query := `
		SELECT
			id,
			title,
			content,
			image,
			created_at,
			COUNT(1) FILTER (WHERE deleted_at IS NULL) OVER () AS total_count
		FROM news
		WHERE deleted_at IS NULL AND created_at > NOW() - INTERVAL '5 hour'
		LIMIT $1 OFFSET $2`
	rows, err := n.db.QueryContext(ctx, query,
		req.Limit,
		req.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r models.News
		err = rows.Scan(
			&r.ID,
			&r.Title,
			&r.Content,
			&r.Image,
			&r.CreatedAt,
			&resp.TotalCount,
		)
		if err != nil {
			return nil, err
		}
		resp.NewsList = append(resp.NewsList, &r)
	}
	return resp, err
}

func (n *newsRepo) GetNewsByID(ctx context.Context, req *models.GetNewsByIDRequest) (
	resp *models.GetNewsByIDResponse, err error) {

	var (
		respID        sql.NullString
		respTitle     sql.NullString
		respContent   sql.NullString
		respImage     sql.NullString
		respCreatedAt sql.NullString
	)
	query := `
		SELECT
			id,
			title,
			content,
			image,
			created_at
		FROM news
		WHERE id = $1`
	err = n.db.QueryRowContext(ctx, query, req.NewsID).Scan(
		&respID,
		&respTitle,
		&respContent,
		&respImage,
		&respCreatedAt,
	)

	resp = &models.GetNewsByIDResponse{
		News: &models.News{
			ID:        respID.String,
			Title:     respTitle.String,
			Content:   respContent.String,
			Image:     respImage.String,
			CreatedAt: respCreatedAt.String,
		},
	}

	if err != nil {
		return nil, err
	}

	return resp, err
}
