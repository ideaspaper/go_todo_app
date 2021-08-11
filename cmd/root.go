package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ideaspaper/puttask/controllers"
	"github.com/ideaspaper/puttask/services"
	"github.com/ideaspaper/puttask/views"
	"github.com/spf13/cobra"
)

var (
	todoService    services.ITodoService
	todoView       views.ITodoView
	todoController controllers.ITodoController
)

func initApp() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exPath := filepath.Dir(ex)
	fileName := filepath.Join(exPath, "task_list.json")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("file %v does not exist\n", fileName)
		fmt.Printf("creating %v\n", fileName)
		os.WriteFile(fileName, []byte("[]"), 0666)
		fmt.Printf("file %v created successfully\n", fileName)
		os.Exit(1)
	}
	todoService, err := services.NewTodoService(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	todoController = controllers.NewTodoController(todoService, views.NewTodoView())
}

var rootCmd = &cobra.Command{
	Use:   "puttask",
	Short: "Puttask is a small todo app on your terminal",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	initApp()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
