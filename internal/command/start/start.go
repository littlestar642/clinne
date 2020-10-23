package start

import (
	"clinne/internal/controller"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "start",
		Short:   "Start the game",
		Long:    "This command is used to start a new game",
		Example: "clinne start",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			err := controller.StartGame()
			if err != nil {
				fmt.Print(err.Error())
			}
		},
	}
}
