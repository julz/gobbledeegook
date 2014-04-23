package gobbledeegook_test

import (
	. "github.com/julz/gobbledeegook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generate", func() {
	Context("When the translation file is empty", func() {
		It("Generates an empty struct based on the language name", func() {
			Expect(Generate("T_EN_US", "{}")).To(Equal("T_EN_US = struct {\n\n}}"))
		})
	})

	Context("When there is a simple message", func() {
		It("Generates a struct with a single, initialized message", func() {
			Expect(Generate("T_EN_US", "{ \"Foo\": \"abc\" }")).To(Equal("T_EN_US = struct {\n\tFoo string\n}{\"abc\"}"))
		})
	})

	Context("When there are multiple messages", func() {
		It("Generates a struct with multiple messages", func() {
			Expect(Generate("T_EN_US", "{ \"Foo\": \"abc\", \"Bar\": \"easy as 123\" }")).To(Equal("T_EN_US = struct {\n\tFoo string\n\tBar string\n}{\"abc\",\"easy as 123\"}"))
		})
	})
})

var _ = Describe("GenerateAll", func() {
	Context("When there are multiple translations", func() {

	})
})
