package terraform

import (
	"bytes"
	"context"
	"os"

	"github.com/clouddrove/smurf/configs"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/pterm/pterm"
)

// Plan runs 'terraform plan' and outputs the plan to the console
func Plan(varNameValue string, varFile string) error {
	tf, err := getTerraform()
	if err != nil {
		return err
	}

	var outputBuffer bytes.Buffer

	customWriter := &configs.CustomColorWriter{
		Buffer: &outputBuffer,
		Writer: os.Stdout,
	}

	tf.SetStdout(customWriter)
	tf.SetStderr(os.Stderr)

	pterm.Info.Println("Running Terraform plan...")
	spinner, _ := pterm.DefaultSpinner.Start("Running terraform plan")

	if varNameValue != "" {
		pterm.Info.Printf("Setting variable: %s\n", varNameValue)
		tf.Plan(context.Background(), tfexec.Var(varNameValue))
	}

	if varFile != "" {
		pterm.Info.Printf("Setting variable file: %s\n", varFile)
		_, err = tf.Plan(context.Background(), tfexec.VarFile(varFile))
		if err != nil {
			spinner.Fail("Terraform plan failed")
			pterm.Error.Printf("Terraform plan failed: %v\n", err)
			return err
		}
	}

	_, err = tf.Plan(context.Background())
	if err != nil {
		spinner.Fail("Terraform plan failed")
		pterm.Error.Printf("Terraform plan failed: %v\n", err)
		return err
	}
	spinner.Success("Terraform plan completed successfully")

	return nil
}
