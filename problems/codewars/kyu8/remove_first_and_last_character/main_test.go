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

var _ = Describe("RemoveChar", func() {
	Context("With country", func() {
		It("Return ountr", func() {
			Expect(RemoveChar("country")).To(Equal("ountr"))
		})
	})

	Context("With person", func() {
		It("Return erso", func() {
			Expect(RemoveChar("person")).To(Equal("erso"))
		})
	})

	Context("With hi", func() {
		It("Return ''", func() {
			Expect(RemoveChar("hi")).To(Equal(""))
		})
	})

	Context("With h", func() {
		It("Return 'h'", func() {
			Expect(RemoveChar("h")).To(Equal(""))
		})
	})
})
