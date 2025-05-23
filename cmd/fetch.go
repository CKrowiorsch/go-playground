package cmd

import (
	"fmt"
	"os"
	"zeiterfassung/zeitfetch"

	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
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
