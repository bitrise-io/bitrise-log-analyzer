package cmd

import (
	"github.com/bitrise-tools/bitrise-log-analyzer/editor"
	"github.com/spf13/cobra"
)

// editorCmd represents the editor command
var editorCmd = &cobra.Command{
	Use:   "editor",
	Short: `Editor`,
	Long:  `Editor`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return editor.LaunchEditor()
	},
}

func init() {
	RootCmd.AddCommand(editorCmd)
}
