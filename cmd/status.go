package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Zeigt den Status der Zeiterfassung",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Status der Zeiterfassung anzeigen.")
		fmt.Println("Status der Zeiterfassung anzeigen.")
	},
}
