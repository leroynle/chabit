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
	"os"
	"strconv"

	"github.com/kyokomi/emoji/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showTasks()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showTasks() {
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
	tasks := convertMapToStruct(nodes)
	data := [][]string{}

	emoji.Println(":ribbon: You got tasks to do today!!! Let's beat them up'")
	fmt.Println()

	for i := range tasks {
		var completionPrint string
		goalPrint := strconv.FormatInt(tasks[i].Complete, 10) + "/" + strconv.FormatInt(tasks[i].Goal, 10)
		if tasks[i].Complete != tasks[i].Goal {
			completionPrint = emoji.Sprint(":cross_mark:")
		} else {
			completionPrint = emoji.Sprint(":check_mark_button:")
		}
		data = append(data, []string{strconv.FormatInt(tasks[i].TaskID, 10), tasks[i].Title, goalPrint, completionPrint})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "TASK", "GOAL", "COMPLETE"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
