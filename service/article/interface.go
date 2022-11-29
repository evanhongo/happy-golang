package service

import "github.com/evanhongo/happy-golang/entity"

//go:generate mockgen -source interface.go -destination ../../mock/mock_service_article.go -package mock

type IArticleService interface {
	GetAllArticles() ([]entity.Article, error)
}
