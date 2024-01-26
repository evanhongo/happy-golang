package repository

import (
	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/pkg/cache"
	"github.com/jinzhu/gorm"
)

type ArticleRepo struct {
	db    *gorm.DB
	cache cache.ICache
}

func (repo *ArticleRepo) GetAllArticles() (res []entity.Article, err error) {
	err = repo.db.Model(&entity.Article{}).
		Select("id,title,content, updated_at, created_at").Find(&res).Error
	return
}

func NewArticleRepo(db *gorm.DB, cache cache.ICache) IArticleRepo {
	return &ArticleRepo{
		db,
		cache,
	}
}
