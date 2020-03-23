/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "relay",
	Short: "A command-line tool to interact with the relay.sh service",
	Long: `This command works with the relay.sh service to provide
a one-stop shop for all your devops workflow automation needs.

To get started, you'll need a relay.sh account - sign up for free
by following this link: 🔗 https://relay.sh/account/signup

Once you've signed up, run this to get started:
▶️  relay auth login
		
You'll probably then want to get started with one of the
default included workflows. 
▶️  relay workflow list
▶️  relay workflow download pagerduty-to-github-example`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	outputformat := "text"

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&outputformat, "out", "", "output format: text, json (default is text)")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "assume 'yes', avoiding confirmation prompts")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "increase verbosity to aid troubleshooting")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable maximum debugging output")
	rootCmd.PersistentFlags().BoolP("color", "c", true, "enable colorized output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
