package stf

import (
	"github.com/clouddrove/smurf/internal/terraform"
	"github.com/spf13/cobra"
)

var destroyApprove bool
var destroyLock bool

// destroyCmd defines a subcommand that destroys the Terraform Infrastructure.
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the Terraform Infrastructure",
	RunE: func(cmd *cobra.Command, args []string) error {
		return terraform.Destroy(destroyApprove, destroyLock)
	},
	Example: `
	smurf stf destroy
	`,
}

func init() {
	destroyCmd.Flags().BoolVar(&destroyApprove, "approve", false, "Skip interactive approval of plan before applying")
	destroyCmd.Flags().BoolVar(&destroyLock, "lock", true, "Don't hold a state lock during the operation. This is dangerous if others might concurrently run commands against the same workspace.")
	stfCmd.AddCommand(destroyCmd)
}
