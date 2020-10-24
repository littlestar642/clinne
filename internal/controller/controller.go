package controller

import (
	"clinne/internal/constants"
	"clinne/internal/model"
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/mbndr/figlet4go"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func StartGame() error {
	fileUtil := file.New()
	files, err := fileUtil.ReadDirectory(constants.MetadataFilePath)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	for {
		chosenVal := rand.Intn(len(files))
		metadata, err := loadFromFile(files[chosenVal])
		if err != nil {
			return err
		}
		codestring, err := fileUtil.ReadFile(metadata.FilePath)
		if err != nil {
			return err
		}
		printer.Println(constants.CodeBlock, color.Underline, color.FgCyan)
		clearSpace()
		printer.Println(codestring, color.FgHiWhite)
		var answer string
		printer.Println("AGREE[Y] OR DISAGREE[N]", color.FgMagenta)
		fmt.Scanln(&answer)
		clearSpace()
		checkAnswer(answer, metadata)
		clearSpace()
		printer.Println("Enter Y to try again! Anything other key to exit", color.FgMagenta)
		fmt.Scanln(&answer)
		if strings.EqualFold(answer, "Y") {
			fmt.Print(constants.ClearScreen)
			continue
		} else {
			break
		}
	}
	return nil
}

func clearSpace() {
	fmt.Println()
	fmt.Println()
}

func checkAnswer(answer string, metadata *model.Metadata) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()

	if (strings.EqualFold(answer, "Y") && metadata.Answer) || (strings.EqualFold(answer, "N") && !metadata.Answer) {
		options.FontColor = []figlet4go.Color{
			figlet4go.ColorYellow,
			figlet4go.ColorCyan,
			figlet4go.ColorGreen,
		}
		printer.Println(constants.SuccessMessage, color.FgGreen)
		renderStr, err := ascii.RenderOpts(constants.NneNinja, options)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(renderStr)
	} else {
		options.FontColor = []figlet4go.Color{
			figlet4go.ColorRed,
			figlet4go.ColorWhite,
			figlet4go.ColorMagenta,
		}
		printer.Println(constants.FailureMessage, color.FgRed)
		renderStr, err := ascii.RenderOpts(constants.DoTheHonours, options)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(renderStr)
	}
	printer.Println("Things to take care:", color.Underline, color.FgHiWhite)
	clearSpace()
	printer.Println(metadata.Rules, color.FgHiYellow)
}

func loadFromFile(filepath string) (*model.Metadata, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	metadata := &model.Metadata{}
	err = json.Unmarshal(byteValue, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
