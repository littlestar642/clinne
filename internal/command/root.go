package command

import (
	"clinne/internal/command/reset"
	"clinne/internal/command/rules"
	"clinne/internal/command/start"
	"clinne/internal/command/stats"
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

	statsCmd := stats.NewCmd()
	if statsCmd != nil {
		rootCmd.AddCommand(statsCmd)
	}

	resetCmd := reset.NewCmd()
	if resetCmd != nil {
		rootCmd.AddCommand(resetCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
