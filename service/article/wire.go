package service

import (
	repository "github.com/evanhongo/happy-golang/repository/article"
	"github.com/google/wire"
)

var ArticleServiceSet = wire.NewSet(
	repository.NewArticleRepo,
	NewArticleService,
)
