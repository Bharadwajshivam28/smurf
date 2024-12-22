package helm

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/clouddrove/smurf/configs"
	"github.com/clouddrove/smurf/internal/helm"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	provisionNamespace string
)

var provisionCmd = &cobra.Command{
	Use:   "provision [RELEASE] [CHART]",
	Short: "Combination of install, upgrade, lint, and template for Helm",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var releaseName, chartPath string

		if len(args) >= 1 {
			releaseName = args[0]
		}
		if len(args) >= 2 {
			chartPath = args[1]
		}

		if releaseName == "" || chartPath == "" {
			data, err := configs.LoadConfig(configs.FileName)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if releaseName == "" {
				releaseName = data.Selm.ReleaseName
				if releaseName == "" {
					releaseName = filepath.Base(data.Selm.ChartName)
				}
			}

			if chartPath == "" {
				chartPath = data.Selm.ChartName
			}

			if releaseName == "" || chartPath == "" {
				return errors.New(color.RedString("RELEASE and CHART must be provided either as arguments or in the config"))
			}

			if provisionNamespace == "" && data.Selm.Namespace != "" {
				provisionNamespace = data.Selm.Namespace
			}
		}

		if releaseName == "" || chartPath == "" {
			return errors.New(color.RedString("RELEASE and CHART must be provided"))
		}

		if provisionNamespace == "" {
			provisionNamespace = "default"
		}

		err := helm.HelmProvision(releaseName, chartPath, provisionNamespace)
		if err != nil {
			return fmt.Errorf(color.RedString("Helm provision failed: %v", err))
		}
		return nil
	},
	Example: `
smurf selm provision my-release ./mychart
smurf selm provision
# In this example, it will read RELEASE and CHART from the config file
smurf selm provision my-release ./mychart -n custom-namespace
`,
}

func init() {
	provisionCmd.Flags().StringVarP(&provisionNamespace, "namespace", "n", "", "Specify the namespace to provision the Helm chart")
	selmCmd.AddCommand(provisionCmd)
}
