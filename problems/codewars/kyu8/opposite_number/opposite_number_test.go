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

var _ = Describe("Opposite", func() {
    Context("With 1", func() {
        It("Return -1", func() {
            Expect(-1).To(Equal(Opposite(1)))
        })
    })
    Context("With -3", func() {
        It("Return 3", func() {
            Expect(3).To(Equal(Opposite(-3)))
        })
    })
    Context("With 0", func() {
        It("Return 0", func() {
            Expect(0).To(Equal(Opposite(0)))
        })
    })
})
