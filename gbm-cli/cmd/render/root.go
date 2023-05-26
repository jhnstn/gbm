package render

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wordpress-mobile/gbm/cmd/render/checklist"
	"github.com/wordpress-mobile/gbm/cmd/render/pr"
)

// renderCmd represents the render command
var RootCmd = &cobra.Command{
	Use:   "render",
	Short: "renders various GBM templates",
	Long: `Use this command to render:
	- release checklists
	- release PRs
	- sync PRs
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("render called")
	},
}

func init() {
	RootCmd.AddCommand(pr.RootCmd)
	RootCmd.AddCommand(checklist.RootCmd)
}
