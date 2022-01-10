package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	Profiles map[string]Profile `yaml:"profiles"`
}

type Profile struct {
	name  string `yaml:"name"`
	email string `yaml:"email"`
}

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "gitswitch",
		Short: "A CLI for switching between multiple git users",
		Long:  `A CLI for switching between multiple git users`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello CLI")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
