package main

import (
	"zeiterfassung/cmd"
)

func main() {
	cmd.RootCmd.AddCommand(cmd.StartCmd, cmd.StopCmd, cmd.StatusCmd, cmd.FetchCmd)
	cmd.Execute()
}
