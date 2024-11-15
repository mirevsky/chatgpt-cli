package cmd_test

import (
	"github.com/mirevsky/chatgpt-cli/cmd"
	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image Command", func() {
	var rootCmd *cobra.Command
	commandName := "image"

	BeforeEach(func() {
		rootCmd = cmd.NewRootCmd()
	})

	It("should find command", func() {
		var thisCmd *cobra.Command
		Ω(rootCmd.Commands()).To(ContainElement(HaveField("Use", commandName), &thisCmd))
		Ω(thisCmd.Name()).To(Equal(commandName))
	})
})
