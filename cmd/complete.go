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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		completeTasks(args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeTasks(argsID string) {
	/*

	 */
	tID, _ := strconv.ParseInt(argsID, 10, 64)

	byteValue, err := ioutil.ReadFile("data/tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatal(err)
	}

	nodes := result["tasks"].([]interface{})
	for _, node := range nodes {
		m := node.(map[string]interface{})
		if id, found := m["TaskID"]; found {
			intID := int64(id.(float64))
			if intID == tID {
				if c, exist := m["Complete"]; exist {
					intC := int64(c.(float64))
					if g, exist := m["Goal"]; exist {
						intG := int64(g.(float64))
						if intC < intG {
							m["Complete"] = intC + 1
						} else {
							fmt.Println("Task #1 is completed!!! Well done")
						}
					}
				}
			}
		}
	}
	byteValue, err = json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("data/tasks.json", byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
