package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-example",
	Short: "Example Go CLI Example Application",
	Long:  `Example Go CLI Example Application`,
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command",
	Long:  "Start command",
	Run: func(cmd *cobra.Command, args []string) {
		isStart, _ := cmd.Flags().GetBool("start")

		if isStart {
			start()
		} else {
			example()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("start", "s", false, "Get Start Function")
	startCmd.Flags().BoolP("example", "e", false, "Get Example Function")
}

func start() {
	fmt.Println("Start Function")
}

func example() {
	fmt.Println("Example Function")
}
