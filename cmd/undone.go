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

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Undone a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid index", err)
		}
		file, err := os.ReadFile("todo.txt")
		if err != nil {
			fmt.Println("error while reading the file", err)
		}

		tasks := strings.Split(strings.TrimSpace(string(file)), "\n")
		if index < 1 || index > len(tasks) {
			fmt.Println("invalid index")
			return
		}

		markUndone := tasks[index-1]
		if !strings.Contains(markUndone, "✅ Completed") {
			fmt.Println("This task is already marked as done.")
			return
		}

		tasks[index-1] = strings.ReplaceAll(tasks[index-1], "✅ Completed: ", "")
		fmt.Println("marking as incomplete: ", tasks[index-1])
		os.WriteFile("todo.txt", []byte(strings.Join(tasks, "\n")), 0644)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
