/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a profile to gitswitch.",
	Long:  `Adds a profile to gitswitch.`,
	Run: func(cmd *cobra.Command, args []string) {
		nameValidate := func(input string) error {
			if input == "" {
				return errors.New("Please enter your git name")
			}
			return nil
		}

		namePrompt := promptui.Prompt{
			Label:    "Name",
			Validate: nameValidate,
		}

		name, err := namePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		emailValidate := func(input string) error {
			_, err := mail.ParseAddress(input)
			if err != nil {
				return errors.New("Please enter a valid email")
			}
			return nil
		}

		emailPrompt := promptui.Prompt{
			Label:    "Email",
			Validate: emailValidate,
		}

		email, err := emailPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		nicknameValidate := func(input string) error {
			if input == "" {
				return errors.New("Please enter a nickname for this git profile")
			}
			return nil
		}

		nicknamePrompt := promptui.Prompt{
			Label:    "Nickname",
			Validate: nicknameValidate,
		}

		nickname, err := nicknamePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", name)
		fmt.Printf("You choose %q\n", email)
		fmt.Printf("You choose %q\n", nickname)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
