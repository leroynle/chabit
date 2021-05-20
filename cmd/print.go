/*
Copyright © Leroy N Le contact@leroynle.com

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
	"chabit/internal"
	"fmt"
	"os"
	"strconv"

	"github.com/kyokomi/emoji/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print your tasks list right on the console",
	Run: func(cmd *cobra.Command, args []string) {
		printTasks()
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printTasks() {

	tasks := internal.GetDataFromJsonTasks()
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
		data = append(data, []string{strconv.FormatInt(tasks[i].TaskID, 10), tasks[i].Title, tasks[i].Duration, goalPrint, completionPrint})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "TASK", "DURATION", "GOAL", "COMPLETE"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
