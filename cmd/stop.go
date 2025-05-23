package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stoppt die Zeiterfassung",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zeiterfassung gestoppt.")
	},
}
