package service

import (
	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/internal/cache"
	repository "github.com/evanhongo/happy-golang/repository/article"
)

type ArticleService struct {
	repo  repository.IArticleRepo
	cache cache.ICache
}

func (service *ArticleService) GetAllArticles() (res []entity.Article, err error) {
	// business logic
	res, err = service.repo.GetAllArticles()
	if err != nil {
		return nil, err
	}
	return
}

func NewArticleService(repo repository.IArticleRepo, cache cache.ICache) IArticleService {
	return &ArticleService{
		repo,
		cache,
	}
}
