/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid index", err)
		}

		file, err := os.ReadFile("todo.txt")
		if err != nil {
			fmt.Println("error reading the file", file)
		}

		tasks := strings.Split(strings.TrimSpace(string(file)), "\n")
		if index < 1 || index > len(tasks) {
			fmt.Println("invalid index")
		}

		fmt.Println("deleted the task: ", tasks[index-1])
		tasks = append(tasks[:index-1], tasks[index:]...)
		os.WriteFile("todo.txt", []byte(strings.Join(tasks, "\n")), 0644)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
