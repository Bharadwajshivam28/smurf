package docker

import (
	"errors"
	"fmt"

	"github.com/clouddrove/smurf/configs"
	"github.com/clouddrove/smurf/internal/docker"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag [SOURCE_IMAGE[:TAG]] [TARGET_IMAGE[:TAG]]",
	Short: "Tag a Docker image for a remote repository",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var source, target string

		if len(args) >= 1 {
			source = args[0]
		}
		if len(args) >= 2 {
			target = args[1]
		}

		if source == "" || target == "" {
			data, err := configs.LoadConfig(configs.FileName)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if source == "" {
				source = data.Sdkr.ImageName
			}
			if target == "" {
				target = data.Sdkr.TargetImageTag
			}

			if source == "" || target == "" {
				return errors.New("both SOURCE and TARGET must be provided either as arguments or in the config")
			}
		}

		pterm.Info.Printf("Tagging image from %q to %q...\n", source, target)
		opts := docker.TagOptions{
			Source: source,
			Target: target,
		}
		if err := docker.TagImage(opts); err != nil {
			return fmt.Errorf("failed to tag image: %w", err)
		}
		pterm.Success.Printf("Successfully tagged image from %q to %q.\n", source, target)
		return nil
	},
	Example: `
  smurf sdkr tag my-app:latest my-org/my-app:prod
  smurf sdkr tag
  # In the second example, it reads SOURCE and TARGET from the config file
`,
}

func init() {
	sdkrCmd.AddCommand(tagCmd)
}
