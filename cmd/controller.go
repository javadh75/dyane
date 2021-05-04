package cmd

import (
	"github.com/javadh75/dyane/controller"
	"github.com/spf13/cobra"
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "This commands run controller.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		controller.Run()
	},
}
