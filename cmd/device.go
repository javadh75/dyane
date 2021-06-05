package cmd

import (
	"errors"

	"github.com/javadh75/dyane/device"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "dev [name]",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a device name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		device.GetLinkCmd(args[0])
	},
}

var linkShowAllCmd = &cobra.Command{
	Use:   "show",
	Short: "List link(s)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		device.GetAllLinksCmd()
	},
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link utilities",
	Long:  ``,
}

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "This commands run device agent.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		device.Run(args)
	},
}

func init() {
	linkShowAllCmd.AddCommand(devCmd)

	linkCmd.AddCommand(linkShowAllCmd)

	deviceCmd.AddCommand(linkCmd)
}
