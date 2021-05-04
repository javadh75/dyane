package cmd

import (
	"github.com/javadh75/dyane/proxy"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "This commands run proxy.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Run()
	},
}
