package cmd

import (
	"github.com/javadh75/dyane/device"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "This commands run device agent.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		device.Run()
	},
}
