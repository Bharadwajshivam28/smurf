package helm

import (
	"github.com/clouddrove/smurf/internal/helm"
	"github.com/spf13/cobra"
)

var createChartCmd = &cobra.Command{
	Use:   "create [NAME] [DIRECTORY]",
	Short: "Create a new Helm chart in the specified directory.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return helm.CreateChart(args[0], args[1])
	},
	Example: `
	smurf helm create mychart ./mychart
	`,
}

func init() {
	selmCmd.AddCommand(createChartCmd)
}
