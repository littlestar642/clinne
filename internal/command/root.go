package command

import (
	"clinne/internal/command/start"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clinne",
	Short: "Play the NNE Game in cli",
	Long:  `Improve on the non-negotiable etiquettes and save yourself from the "Doing-The-Honours"`,
}

// Execute function executes the root command of clinne
func Execute() error {
	startCmd := start.NewCmd()
	if startCmd != nil {
		rootCmd.AddCommand(startCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
