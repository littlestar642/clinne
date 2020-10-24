package reset

import (
	"clinne/internal/constants"
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "reset",
		Short:   "Reset your game stats",
		Long:    "This command is used to reset the game and start from fresh",
		Example: "clinne reset",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fileUtil := file.New()
			isExist, err := fileUtil.IsFileExist(constants.ResultFilePath)
			if err != nil {
				printer.Println(fmt.Sprintf("error in reading file %s", err.Error()), color.FgRed)
			}
			if !isExist {
				printer.Println("Seems like you have not played any game yet. Play one to generate stats!")
			} else {
				fileUtil.DeleteFile(constants.ResultFilePath)
				printer.Println("Your game has been refreshed!!", color.FgMagenta)
			}
		},
	}
}
