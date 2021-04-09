package cmd

import (
	"github.com/spf13/cobra"
)

// RecErrorCmd represents the RecError command
var RecErrorCmd = &cobra.Command{
	Use:   "rec-error",
	Short: "Recording error notification command",
	Long: `This command notifies you of recording errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		Slack(" :x: ","warning")
	},
}

func init() {
	rootCmd.AddCommand(RecErrorCmd)
}
