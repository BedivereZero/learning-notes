package ginkgo_test

import (
	"github.com/BedivereZero/learning-notes/golang/ginkgo/books"
	"github.com/go-logr/zapr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"k8s.io/klog/v2/textlogger"
)

func init() {
	GinkgoLogr = zapr.NewLogger(zap.Must(zap.NewDevelopment()))
	GinkgoLogr = textlogger.NewLogger(textlogger.NewConfig())
}

var _ = Describe("Books", func() {
	var book *books.Book

	BeforeEach(func() {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		Expect(book.IsValid()).To(BeTrue())
		GinkgoLogr.Info("set book", "book", book)
	})

	Describe("Extracting the author's first and last name", func() {
		Context("When the author only has one name", func() {
			BeforeEach(func() {
				book.Author = "Hugo"
			})

			It("interprets the single author name as a last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("returns empty for the first name", func() {
				Expect(book.AuthorFirstName()).To(BeEmpty())
			})
		})

		Context("When the author has both names", func() {
			It("can extract the author's last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("can extract the author's first name", func() {
				Expect(book.AuthorFirstName()).To(Equal("Victor"))
			})
		})

	})
})
