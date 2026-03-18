package cmd

import (
	"fmt"
	"log"
	"os"
	"zeiterfassung/zeitfetch"

	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:   "fetch <url>",
	Short: "Liest die aktuelle Zeit von einer Webseite aus",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.Printf("Starte Zeitabfrage für URL: %s", args[0])
		}
		time, err := zeitfetch.FetchTimeFromWeb(args[0])
		if err != nil {
			if verbose {
				log.Printf("Fehler beim Auslesen der Zeit: %v", err)
			}
			os.Exit(1)
		}
		if verbose {
			log.Printf("Gefundene Zeit: %s", time)
		}
		fmt.Println("Gefundene Zeit:", time)
	},
}

func init() {
	FetchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Aktiviere ausführliche Logausgabe")
}
