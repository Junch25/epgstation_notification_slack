package cmd

import (
	"github.com/spf13/cobra"
)

// RecEndCmd represents the RecEnd command
var RecEndCmd = &cobra.Command{
	Use:   "rec-end",
	Short: "Recording end notification command",
	Long: `This command notifies the end of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		Slack(" :white_check_mark: ","good")
	},
}

func init() {
	rootCmd.AddCommand(RecEndCmd)
}
