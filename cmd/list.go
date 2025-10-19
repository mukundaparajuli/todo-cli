/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all the todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		file, error := os.Open("todo.txt")
		if error != nil {
			fmt.Println("error listing the tasks", error)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		i := 1
		for scanner.Scan() {
			fmt.Printf("%d. %s \n", i, scanner.Text())
			i++
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
