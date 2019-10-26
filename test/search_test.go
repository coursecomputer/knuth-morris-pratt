package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AlgoCafe/knuth-morris-pratt/kmp"
)

var _ = Describe("Search", func() {
	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abcdabd\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabd")

			// Process
			index = kmp.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(15))
		})

		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abcdabdr\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabdr")

			// Process
			index = kmp.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(-1))
		})

		It("with \"abc abcdab abcdabcdabde\" as buffer and \"de\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("de")

			// Process
			index = kmp.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(21))
		})
	})
})