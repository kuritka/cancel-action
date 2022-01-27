package cmd

import (
	"github.com/kuritka/cancel-action/internal/common"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "version",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("%s %s %s", BrightWhite("version:"), BrightWhite("to be implemented. see: "), common.HomeURL)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
