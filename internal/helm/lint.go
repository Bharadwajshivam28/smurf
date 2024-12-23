package helm

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/strvals"
)

func HelmLint(chartPath string, fileValues []string) error {
	spinner, _ := pterm.DefaultSpinner.Start("Linting chart")
	defer spinner.Stop()

	client := action.NewLint()

	vals := make(map[string]interface{})
	for _, f := range fileValues {
		additionalVals, err := chartutil.ReadValuesFile(f)
		if err != nil {
			color.Red("Failed to read values file '%s': %v \n", f, err)
			return err
		}
		for key, value := range additionalVals {
			vals[key] = value
		}
	}

	for _, set := range fileValues {
		if err := strvals.ParseInto(set, vals); err != nil {
			color.Red("Failed to parse set values '%s': %v \n", set, err)
			return err
		}
	}

	result := client.Run([]string{chartPath}, vals)
	if len(result.Messages) > 0 {
		for _, msg := range result.Messages {
			color.Yellow("Severity: %s \n", msg.Severity)
			color.Yellow("Path: %s \n", msg.Path)
			fmt.Println(msg)
			fmt.Println()
		}
		spinner.Fail("Linting issues found \n")
	} else {
		color.Green("No linting issues found in the chart %s \n", chartPath)
		spinner.Success("Linting completed successfully \n")
	}
	return nil
}
