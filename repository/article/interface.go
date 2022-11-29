package repository

import "github.com/evanhongo/happy-golang/entity"

//go:generate mockgen -source interface.go -destination ../../mock/mock_repository_article.go -package mock

type IArticleRepo interface {
	GetAllArticles() ([]entity.Article, error)
}
