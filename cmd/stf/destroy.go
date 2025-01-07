package stf

import (
	"github.com/clouddrove/smurf/internal/terraform"
	"github.com/spf13/cobra"
)

// destroyCmd defines a subcommand that destroys the Terraform Infrastructure.
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the Terraform Infrastructure",
	RunE: func(cmd *cobra.Command, args []string) error {
		return terraform.Destroy(approve)
	},
	Example: `
	smurf stf destroy
	`,
}

func init() {
	destroyCmd.Flags().BoolVar(&approve, "approve", false, "Skip interactive approval of plan before applying")
	stfCmd.AddCommand(destroyCmd)
}
