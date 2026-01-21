package cmd

import (
	"github.com/spf13/cobra"
	"github.com/winnerx0/reqx/internal/command"
)

var rootCmd = &cobra.Command{
	Use: "reqx",
	Short: "Reqx http client",
	Long: "This is an application that acts like a http client with json files",
}

func Execute(){
	cobra.CheckErr(rootCmd.Execute())
}

func init(){
	rootCmd.AddCommand(command.SendCmd)
}