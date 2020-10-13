package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("PositiveSum", func() {
	Context("With []int{1, 2}", func() {
		It("Return 3", func() {
			Expect(PositiveSum([]int{1, 2})).To(Equal(3))
		})
	})

	Context("With []int{1, 2, -4}", func() {
		It("Return 3", func() {
			Expect(PositiveSum([]int{1, 2, -4})).To(Equal(3))
		})
	})

	Context("With []int{}", func() {
		It("Return 0", func() {
			Expect(PositiveSum([]int{})).To(Equal(0))
		})
	})
})
