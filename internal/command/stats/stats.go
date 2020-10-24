package stats

import (
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "stats",
		Short:   "Stats of your game",
		Long:    "This command is used to show your stats for clearing NNEs",
		Example: "clinne stats",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fileUtil := file.New()
			isExist, err := fileUtil.IsFileExist("/results/result.txt")
			if err != nil {
				printer.Println(fmt.Sprintf("error in reading file %s", err.Error()), color.FgRed)
			}
			if !isExist {
				printer.Println("Seems like you have not played any game yet. Play one to generate stats!")
			}
		},
	}
}
