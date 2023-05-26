package pr

import (
	"fmt"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var RootCmd = &cobra.Command{
	Use:   "pr",
	Short: "render the body for a release PR",
	Long: `
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("render pr called")
	},
}

func init() {
	RootCmd.AddCommand(releaseCmd)
	RootCmd.AddCommand(syncCmd)
}
