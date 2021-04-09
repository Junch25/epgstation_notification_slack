package cmd

import (
	"github.com/spf13/cobra"
)

// RecStartCmd represents the RecStart command
var RecStartCmd = &cobra.Command{
	Use:   "rec-start",
	Short: "Recording start notification command",
	Long: `This command notifies the start of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		Slack(" :arrow_forward: ","danger")
	},
}

func init() {
	rootCmd.AddCommand(RecStartCmd)
}
