/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/clouddrove/smurf/cmd"
	_ "github.com/clouddrove/smurf/cmd/terraform"
)

func main() {
	cmd.Execute()
}
