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

var _ = Describe("Method EvenOrOdd", func() {
	Context("With 1", func() {
		It("Return 'Odd'", func() {
			Expect("Odd").To(Equal(EvenOrOdd(1)))
		})
	})
	Context("With 2", func() {
		It("Return 'Even'", func() {
			Expect("Even").To(Equal(EvenOrOdd(2)))
		})
	})
	Context("With -1", func() {
		It("Return 'Odd'", func() {
			Expect("Odd").To(Equal(EvenOrOdd(-1)))
		})
	})
	Context("With -2", func() {
		It("Return 'Even'", func() {
			Expect("Even").To(Equal(EvenOrOdd(-2)))
		})
	})
})
