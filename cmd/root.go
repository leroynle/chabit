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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chabit",
	Short: "An awesome habits and tasks tracker for your console",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	resetDailyTasks()
	// resetWeeklyTask()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chabit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".chabit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".chabit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func resetDailyTasks() {
	tasks := internal.GetDataFromJsonTasks()

	if isNewDay() {
		for i := range tasks {
			p := &tasks[i]
			if p.Duration == "Daily" {
				p.Complete = 0
			}
		}
		byteValue, err := json.MarshalIndent(tasks, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("data/tasks.json", byteValue, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func isNewDay() bool {
	today := time.Now().Weekday().String()

	getDay := internal.GetDataFromJsonUtilites()
	yesterday := getDay["Today"]
	fmt.Println(yesterday)
	if today != yesterday {
		getDay["Today"] = today
		internal.WriteDataToJsonUtilites(getDay)
		return true
	} else {
		return false
	}
}

// func resetWeeklyTask() {
// 	today := time.Now().Weekday().String()
// 	byteValue, err := ioutil.ReadFile("data/tasks.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var result map[string]interface{}
// 	err = json.Unmarshal(byteValue, &result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	nodes := result["tasks"].([]interface{})

// 	for _, node := range nodes {
// 		m := node.(map[string]interface{})
// 		if d, found := m["Duration"]; found {
// 			if d == "Weekly" {
// 				if w, found1 := m["WeekStart"]; found1 {
// 					if w == today {
// 						m["Complete"] = 0
// 					}
// 				}
// 			}
// 		}
// 	}
// 	byteValue, err = json.MarshalIndent(result, "", " ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = ioutil.WriteFile("data/tasks.json", byteValue, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
