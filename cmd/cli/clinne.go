package main

import (
	"clinne/internal/command"
	"clinne/internal/constants"
	"fmt"
	"github.com/mbndr/figlet4go"
)

func main() {
	fmt.Print(constants.ClearScreen)
	fmt.Print(constants.HomePosition)
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
	}
	renderStr, err := ascii.RenderOpts(constants.OpeningBanner, options)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(renderStr)
	err = command.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
