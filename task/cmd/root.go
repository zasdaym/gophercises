package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root cmd.
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
