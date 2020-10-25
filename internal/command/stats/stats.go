package stats

import (
	"clinne/internal/constants"
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"fmt"
	"github.com/fatih/color"
	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
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
			isExist, _ := fileUtil.IsFileExist("/results/result.txt")
			if !isExist {
				printer.Println("Seems like you have not played any game yet. Play one to generate stats!", color.FgHiGreen)
				return
			}
			fileContent, err := fileUtil.ReadFile(constants.ResultFilePath)
			if err != nil {
				printer.Println(fmt.Sprintf("error in reading file %s", err.Error()), color.FgRed)
				return
			}
			series := strings.Split(fileContent, " ")
			var floatSeries []float64
			for _, val := range series {
				if val == "" {
					continue
				}
				floatVal, err := strconv.ParseFloat(val, 64)
				if err != nil {
					printer.Println(fmt.Sprintf("error in parsing to float %s", err.Error()), color.FgRed)
					break
				}
				floatSeries = append(floatSeries, floatVal)
			}
			printer.Println("Following is a graph that depicts your NNE performance!!", color.Underline, color.FgHiGreen)
			fmt.Println()
			fmt.Println()
			data := floatSeries
			graph := asciigraph.Plot(data, asciigraph.Caption("Game Stats"), asciigraph.Width(50))
			fmt.Println(graph)
		},
	}
}
