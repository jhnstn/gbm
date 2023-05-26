package checklist

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wordpress-mobile/gbm/pkg/render"
)

var Version string

// checklistCmd represents the checklist command
var RootCmd = &cobra.Command{
	Use:   "checklist",
	Short: "render the content for the release checklist",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {

		jsonData := fmt.Sprintf(`
		{
			"version": "%s",
			"scheduled": true,
			"date": "Jan, 01, 1970",
			"message" : "THIS IS A MESSAGE",
			"releaseUrl": "",
			"mobileVersion": ""
		}
		`,
			Version)
		result, err := render.Template("templates/checklist.html", jsonData)

		if err != nil {
			panic(err)
		}
		fmt.Print(err)

		fmt.Print(result)
	},
}

func init() {
	RootCmd.Flags().StringVarP(&Version, "version", "v", "", "release version")

}
