package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var RootCmd = &cobra.Command{
	Use:   "zeiterfassung",
	Short: "Zeiterfassung CLI",
	Long:  `Ein CLI-Tool zur Zeiterfassung und Auswertung von Webseiten-Zeiten.`,
}

func init() {
	RootCmd.AddCommand(StartCmd, StopCmd, StatusCmd, FetchCmd)
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Aktiviere ausführliche Logausgabe")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if verbose {
			log.Printf("Fehler beim Ausführen des RootCmd: %v", err)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
