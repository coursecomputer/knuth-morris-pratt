package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AlgoCafe/knuth-morris-pratt/kmp"
)

var _ = Describe("Table", func() {
	var table = kmp.Table{}

	Context("Success", func() {
		It("with \"participate in parachute\" as argument", func() {
			// Prepare
			var pattern = []byte("participate in parachute")

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal([]int{
				-1, 0, 0, 0, 0, 0, 0, -1, 0, 2, 0, 0, 0, 0, 0, -1, 0, 0, 3, 0, 0, 0, 0, 0, 0,
			}))
		})

		It("with \"abcdabd\" as argument", func() {
			// Prepare
			var pattern = []byte("abcdabd")

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal([]int{
				-1, 0, 0, 0, -1, 0, 2, 0,
			}))
		})

		It("with \"abacababc\" as argument", func() {
			// Prepare
			var pattern = []byte("abacababc")

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal([]int{
				-1, 0, -1, 1, -1, 0, -1, 3, 2, 0,
			}))
		})

		It("with \"abacababa\" as argument", func() {
			// Prepare
			var pattern = []byte("abacababa")

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal([]int{
				-1, 0, -1, 1, -1, 0, -1, 3, -1, 3,
			}))
		})
	})
})