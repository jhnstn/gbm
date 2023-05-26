package pr

import (
	"fmt"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "",
	Long: `
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release called")
	},
}

func init() {
}
