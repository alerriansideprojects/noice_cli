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
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cherreExavaultCmd = &cobra.Command{
	Use: "exavault",
	Short: "Generate a Cherre Exavault Ingest template",
	Long: `This command generates an empty template for a Cherre Exavault Ingest`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Generating Cherre Exavault Ingest\n")
		
		exaType, _ := cmd.Flags().GetString("type")
		projName, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")

		if checkInputTypes(exaType) && checkName(projName) {
			if path != ""{
				generateTemplate(exaType, projName, path)
			} else {
				generateTemplate(exaType, projName, "tmp")
			}
		} else {
			fmt.Println("Missing Flag(s)")
		}
	},
}

func init() {
	generateCmd.AddCommand(cherreExavaultCmd)

	cherreExavaultCmd.Flags().String("name", "", "Generate template with supplied name. (Ex: noice_cli)")
	cherreExavaultCmd.Flags().String("type", "", "Generate template with supplied type. (Ex: csv, json, xml, zip)")
}

func checkInputTypes(exaType string) bool {
	switch exaType {
	case
		"csv", 
		"xml", 
		"json",
		"zip":
		return true
	}
	fmt.Println("Invalid selection for type. Select from 'csv', 'xml', 'json' or 'zip'")
	return false
}

func checkName(name string) bool {
	if name != "" {
		return true
	}
	
	fmt.Println("Must supply name with flag --name=<name>")
	return false
}

func generateTemplate(exaType string, projName string, path string){
	err := os.Mkdir(path, 0755)

	if err != nil {
		log.Fatal(err)
	} else {
		makeStructure(path, projName)
	}

	fmt.Printf("Generated %v template.\nGrab them in the 'tmp' folder!", exaType)
}

func makeStructure(path string, projName string) {
	folderPaths := []string{
		fmt.Sprintf("%v/%v/services/taps/", path, projName),
		fmt.Sprintf("%v/tests", path),
	}
	filePaths := []string{
		fmt.Sprintf("%v/Dockerfile", path),
		fmt.Sprintf("%v/envs", path),
		fmt.Sprintf("%v/mypy.ini", path),
		fmt.Sprintf("%v/pyproject.toml", path),
		fmt.Sprintf("%v/README.md", path),
		fmt.Sprintf("%v/.env", path),
		fmt.Sprintf("%v/mashey-cherre.json", path),
		fmt.Sprintf("%v/test.json", path),
	}

	for _, folderPath := range folderPaths {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, filePath := range filePaths {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		file.Close()
	}
}