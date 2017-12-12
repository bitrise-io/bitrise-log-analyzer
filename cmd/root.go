// Copyright © 2016 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

//var db string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bitrise-log-analyzer",
	Short: "Bitrise Log Analyzer tool",
	Long:  ``,
	//	PersistentPreRun: func(cmd *Command, args []string)
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bitrise-log-analyzer.yaml)")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
