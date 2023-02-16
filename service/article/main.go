package service

import (
	"github.com/evanhongo/happy-golang/entity"
	repository "github.com/evanhongo/happy-golang/repository/article"
)

type ArticleService struct {
	repo repository.IArticleRepo
}

func (service *ArticleService) GetAllArticles() (res []entity.Article, err error) {
	// business logic
	res, err = service.repo.GetAllArticles()
	if err != nil {
		return nil, err
	}
	return
}

func NewArticleService(repo repository.IArticleRepo) IArticleService {
	return &ArticleService{
		repo,
	}
}
