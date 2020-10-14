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

var _ = Describe("RepeatStr", func() {
	Context("With 5, a", func() {
		It("Return aaaaa", func() {
			Expect(RepeatStr(5, "a")).To(Equal("aaaaa"))
		})
	})

	Context("With 3, vb", func() {
		It("Return vbvbvb", func() {
			Expect(RepeatStr(3, "vb")).To(Equal("vbvbvb"))
		})
	})
})
