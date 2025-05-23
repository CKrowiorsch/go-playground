package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Startet die Zeiterfassung",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zeiterfassung gestartet.")
	},
}
