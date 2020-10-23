package create

import (
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCmd(fileUtil file.File) *cobra.Command {
	var metadataFilePath string

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create new code block",
		Long:    "This command helps create new code block by providing new metadata.json",
		Example: "clinne create --path <filepath>/metadata.json",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			printer.Println("Creating a new code block.", color.FgMagenta)
			_, err := fileUtil.GetIoReader(metadataFilePath)
			if err != nil {
				printer.Println(fmt.Sprintf("error in reading file %v", err.Error()), color.FgRed)
				return
			}

		},
	}
	createCmd.Flags().StringVarP(&metadataFilePath, "path", "", "", "clinne create --path <filepath>/metadata.json")
	err := createCmd.MarkFlagRequired("path")
	if err != nil {
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	return createCmd
}
