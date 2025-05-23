package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Startet die Zeiterfassung",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Zeiterfassung gestartet.")
		fmt.Println("Zeiterfassung gestartet.")
	},
}
