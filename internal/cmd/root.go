package cmd

import (
	"os"

	"github.com/kuritka/cancel-action/internal/common"
	"github.com/kuritka/cancel-action/internal/common/log"
	"github.com/spf13/cobra"

	w "github.com/logrusorgru/aurora"
)

var (
	// Verbose output
	Verbose bool
)

var logger = log.Log

var rootCmd = &cobra.Command{
	Short: common.Action,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("%s %s 🛌 🤺", w.BrightMagenta(common.Action), w.BrightYellow("started"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Info().Msgf("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msgf("Not sure what to do next? %s %s", w.BrightWhite("see:"), common.HomeURL)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute runs concrete command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
