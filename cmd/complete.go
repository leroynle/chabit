/*
Copyright © 2021 Leroy N Le contact@leroynle.com

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
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <TaskID>",
	Short: "Mark your tasks completed",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		times, _ := cmd.Flags().GetInt64("times")
		if times > 1 {
			completeTasksWithTimes(args[0], times)
		} else if times < 1 {
			fmt.Println("Number of times must be greater than 0 - For example: chabit complete <TaskID> -t 2")
		} else {
			completeTasks(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	var times int64
	completeCmd.Flags().Int64VarP(&times, "times", "t", 1, "number of times you completed your task")
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

	tasks := internal.GetDataFromJsonTasks()

	for i := range tasks {
		p := &tasks[i]
		intID := int64(p.TaskID)
		if intID == tID {
			intC := int64(p.Complete)
			intG := int64(p.Goal)
			if intC < intG {
				p.Complete = intC + 1
			} else {
				fmt.Println("Task #1 is completed!!! Well done")
			}
		}

	}
	internal.WriteDataToJsonTasks(tasks)
}

func completeTasksWithTimes(argsID string, t int64) {
	/*

	 */
	tID, _ := strconv.ParseInt(argsID, 10, 64)

	tasks := internal.GetDataFromJsonTasks()

	for i := range tasks {
		p := &tasks[i]
		intID := int64(p.TaskID)
		if intID == tID {
			intC := int64(p.Complete)
			intG := int64(p.Goal)
			if intC < intG {
				p.Complete = intC + t
			} else {
				fmt.Println("Task #1 is completed!!! Well done")
			}
		}

	}
	internal.WriteDataToJsonTasks(tasks)
}
