package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stoppt die Zeiterfassung",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Zeiterfassung gestoppt.")
		fmt.Println("Zeiterfassung gestoppt.")
	},
}
