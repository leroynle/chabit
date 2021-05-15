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
	"chabit/internal"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add your new habit or task",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		manual, _ := cmd.Flags().GetBool("manual")
		if manual {
			addTaskManual()
		} else {
			addATask()
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	var manual bool
	addCmd.Flags().BoolVarP(&manual, "manual", "m", false, "manual add a task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addATask() {
	prompt := promptui.Select{
		Label: "Select one option",
		Items: []string{"Create your own", "Choose a preset"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Create your own" {
		addTaskManual()
	} else {
		fmt.Println("Add Later!!!")
	}

	fmt.Printf("You choose %q\n", result)
}

func addTaskManual() {
	weekStart := "Sunday"
	taskTitle := getTaskTitleFromUser()
	taskDuration := getTaskDurationFromUser()
	if taskDuration == "Weekly" {
		weekStart = getWeekStarForWeeklyTask()
	}
	taskGoal := getGoalFromUser(taskDuration)

	tasks := internal.GetDataFromJsonTasks()
	//get max Task ID
	maxTaskID := getMaxTaskID() + 1

	newTask := internal.Tasks{
		Title:     taskTitle,
		TaskID:    maxTaskID,
		Duration:  taskDuration,
		WeekStart: weekStart,
		Goal:      taskGoal,
		Complete:  0,
	}
	tasks = append(tasks, newTask)
	internal.WriteDataToJsonTasks(tasks)
}

func getTaskTitleFromUser() string {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Title",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		log.Fatal(err)
	}
	return result
}

func getTaskDurationFromUser() string {
	prompt1 := promptui.Select{
		Label: "Duration: Determines the period of time for a single completion",
		Items: []string{"Daily", "Weekly"},
	}

	_, result, err := prompt1.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		log.Fatal(err)
	}
	return result
}

func getWeekStarForWeeklyTask() string {
	prompt := promptui.Select{
		Label: "Start Week On",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		log.Fatal(err)
	}
	return result
}

func getGoalFromUser(duration string) int64 {
	validate := func(input string) error {
		num, err := strconv.ParseInt(input, 10, 64)
		if err != nil || num < 0 {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt1 := promptui.Prompt{
		Label:    "Set your Goal",
		Validate: validate,
	}

	result, err := prompt1.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		log.Fatal(err)
	}
	resultInt, _ := strconv.ParseInt(result, 10, 64)
	return resultInt
}

func getMaxTaskID() int64 {
	/*
		Function to get max TaskID from JSON file

		return: the maximum number of taskID from JSON file
	*/
	task := internal.GetDataFromJsonTasks()
	var maxTID int64
	if len(task) == 0 {
		maxTID = 0
	} else {
		for i := range task {
			maxTID = max(maxTID, task[i].TaskID)
		}
	}
	return maxTID
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
