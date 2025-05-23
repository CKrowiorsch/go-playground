package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "zeiterfassung",
	Short: "Zeiterfassung CLI",
	Long:  `Ein CLI-Tool zur Zeiterfassung und Auswertung von Webseiten-Zeiten.`,
}

func init() {
	RootCmd.AddCommand(StartCmd, StopCmd, StatusCmd, FetchCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
