package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a selected password from the store",
	Long: `
- Deletes a selected password from the store`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		passwordName := args[0]
		delete(passwordName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(passwordName string) {
	homeDir, _ := os.UserHomeDir()
	path := homeDir + "/.store/" + passwordName + ".asc"
	cmd := exec.Command("rm", path)

	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("Password successfully removed")
}