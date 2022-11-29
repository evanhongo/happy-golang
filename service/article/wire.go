//go:build wireinject
// +build wireinject

package service

import (
	"github.com/evanhongo/happy-golang/internal/cache"
	repository "github.com/evanhongo/happy-golang/repository/article"
	"github.com/google/wire"
)

var ArticleServiceSet = wire.NewSet(
	repository.NewArticleRepo,
	cache.NewCache,
	NewArticleService,
)
