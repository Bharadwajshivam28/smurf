package helm

import (
	"github.com/clouddrove/smurf/internal/helm"
	"github.com/spf13/cobra"
)

var uninstallNamespace string

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [NAME]",
	Short: "Uninstall a Helm release.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		releaseName := args[0]
		if uninstallNamespace == "" { // If no namespace is provided, use default
            uninstallNamespace = "default"
        }
		return helm.HelmUninstall(releaseName, uninstallNamespace) // Assuming 'default' namespace or make it a flag
	},
}

func init() {
	uninstallCmd.Flags().StringVarP(&uninstallNamespace, "namespace", "n", "", "Specify the namespace to install the Helm chart into")
	selmCmd.AddCommand(uninstallCmd)
}
