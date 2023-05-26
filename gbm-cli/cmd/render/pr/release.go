package pr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "render the body for a release PR",
	Long: `Use to render the body for release PRs that target:
	- WordPress/Gutenberg
	- wordpress-mobile/gutenberg-mobile
	- wordpress-mobile/WordPress-Android
	- wordpress-mobile/WordPress-iOS
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release called")
	},
}

func init() {
	releaseCmd.Flags().StringVarP(&Version, "version", "v", "", "release version")
}
