package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestArticle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Article Suite")
}
