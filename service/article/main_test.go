package service_test

import (
	"errors"
	"time"

	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/mock"
	articleService "github.com/evanhongo/happy-golang/service/article"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArticleService", func() {
	var (
		ctl       *gomock.Controller
		mockRepo  *mock.MockIArticleRepo
		mockCache *mock.MockICache
		service   articleService.IArticleService
	)

	BeforeEach(func() {
		ctl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockIArticleRepo(ctl)
		mockCache = mock.NewMockICache(ctl)
		service = articleService.NewArticleService(mockRepo, mockCache)
	})

	It("should get all articles", func(ctx SpecContext) {
		articles := []entity.Article{{Id: 1, Title: "Hello", Content: "World", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
		gomock.InOrder(
			mockRepo.EXPECT().GetAllArticles().Return(articles, nil),
		)
		res, err := service.GetAllArticles()
		Expect(res).To(Equal(articles))
		Expect(err).To(BeNil())
	})

	It("should not get all articles when error happens in repository", func(ctx SpecContext) {
		err := errors.New("something wrong")
		gomock.InOrder(
			mockRepo.EXPECT().GetAllArticles().Return(nil, err),
		)
		res, err := service.GetAllArticles()
		Expect(res).To(BeNil())
		Expect(err).To(MatchError(err))
	})
})
