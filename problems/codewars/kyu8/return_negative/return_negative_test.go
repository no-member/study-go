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

var _ = Describe("MakeNegative", func() {
	Context("With 1", func() {
		It("Return -1", func() {
			Expect(MakeNegative(1)).To(Equal(-1))
		})
	})

	Context("With -4", func() {
		It("Return -4", func() {
			Expect(MakeNegative(-4)).To(Equal(-4))
		})
	})

	Context("With 0", func() {
		It("Return 0", func() {
			Expect(MakeNegative(0)).To(Equal(0))
		})
	})
})
