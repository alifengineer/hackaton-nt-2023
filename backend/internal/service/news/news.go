package news

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
)

func (s *Service) CreateNews(ctx context.Context, req *models.CreateNewsRequest) (resp *models.News, err error) {

	resp, err = s.strg.News().CreateNews(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetAllNews(ctx context.Context, req *models.GetAllNewRequest) (resp *models.GetAllNewsResponse, err error) {

	resp, err = s.strg.News().GetAllNews(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetNewsByID(ctx context.Context, req *models.GetNewsByIDRequest) (
	resp *models.GetNewsByIDResponse, err error) {

	resp, err = s.strg.News().GetNewsByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}
