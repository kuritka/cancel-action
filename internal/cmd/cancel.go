package cmd

import (
	"github.com/AbsaOSS/env-binder/env"
	"github.com/kuritka/cancel-action/internal/common/runner"
	"github.com/kuritka/cancel-action/internal/impl"
	"github.com/kuritka/cancel-action/internal/impl/cancel"
	"github.com/rs/zerolog"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var cancelCommand = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel workflow",
	Long:  "With respect to the configuration inside the action, it cancel redundant workflow run.",

	Run: func(cmd *cobra.Command, args []string) {
		opts := &impl.ActionOpts{}
		err := env.Bind(opts)
		kingpin.FatalIfError(err, "reading environment variables")
		if opts.Verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		logger.Debug().Msgf("loaded configuration: \n %v", aurora.BrightWhite(opts))
		runner.NewCommonRunner(cancel.NewCommand(*opts)).MustRun()
	},
}

func init() {
	rootCmd.AddCommand(cancelCommand)
}
