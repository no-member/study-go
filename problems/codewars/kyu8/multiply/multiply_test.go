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

var _ = Describe("Multiply", func() {
	Context("With 1 and 1", func() {
		It("Return 1", func() {
			Expect(Multiply(1, 1)).To(Equal(1))
		})
	})
	Context("With 100 and 0", func() {
		It("Return 0", func() {
			Expect(Multiply(100, 0)).To(Equal(0))
		})
	})
})
