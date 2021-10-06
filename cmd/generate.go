/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Aliases: []string{"gen", "g"},
	Short: "Generates an empty project at a specified location",
	Long: `Adds an entire Cherre project, depending on what flags are passed, to a specific location`,
}

var cherreApiCmd = &cobra.Command{
	Use: "api",
	Short: "Generate a Cherre API ingest template.",
	Long: `This command generates an empty template for a Cherre API Ingest`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Generating Cherre API Ingest")
	},
}

var cherreS3Cmd = &cobra.Command{
	Use: "awss3",
	Short: "Generate a Cherre Exavault Ingest template",
	Long: `This command generates an empty template for a Cherre AWS S3 Ingest`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Generating Cherre AWS S3 Ingest")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.AddCommand(cherreApiCmd)
	generateCmd.AddCommand(cherreS3Cmd)
}