package command

import (
	"clinne/internal/command/rules"
	"clinne/internal/command/start"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clinne",
	Short: "Play the NNE Game in cli",
	Long:  `Improve your Non-Negotiable Etiquettes and save yourself from "Doing-The-Honours"`,
}

// Execute function executes the root command of clinne
func Execute() error {
	startCmd := start.NewCmd()
	if startCmd != nil {
		rootCmd.AddCommand(startCmd)
	}

	rulesCmd := rules.NewCmd()
	if rulesCmd != nil {
		rootCmd.AddCommand(rulesCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
