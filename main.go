package main

import (
	"fmt"
	"os"
	"zeiterfassung/zeitfetch"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "zeiterfassung",
		Short: "Zeiterfassung CLI",
		Long:  `Ein CLI-Tool zur Zeiterfassung und Auswertung von Webseiten-Zeiten.`,
	}

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Startet die Zeiterfassung",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Zeiterfassung gestartet.")
		},
	}

	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stoppt die Zeiterfassung",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Zeiterfassung gestoppt.")
		},
	}

	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Zeigt den Status der Zeiterfassung",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Status der Zeiterfassung anzeigen.")
		},
	}

	var fetchCmd = &cobra.Command{
		Use:   "fetch <url>",
		Short: "Liest die aktuelle Zeit von einer Webseite aus",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			time, err := zeitfetch.FetchTimeFromWeb(args[0])
			if err != nil {
				fmt.Println("Fehler beim Auslesen der Zeit:", err)
				os.Exit(1)
			}
			fmt.Println("Gefundene Zeit:", time)
		},
	}

	rootCmd.AddCommand(startCmd, stopCmd, statusCmd, fetchCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
