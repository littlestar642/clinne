package controller

import (
	"clinne/internal/constants"
	"clinne/internal/model"
	"clinne/internal/pkg/file"
	"clinne/internal/pkg/printer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mbndr/figlet4go"
)

func StartGame() error {
	fileUtil := file.New()
	files, err := fileUtil.ReadDirectory(constants.MetadataFilePath)
	if err != nil {
		return err
	}
	err = setUpFile(fileUtil)
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
		printer.Println(metadata.Description+". "+constants.CodeBlock, color.Underline, color.FgCyan)
		clearSpace()
		printer.Println(codestring, color.FgHiWhite)
		var answer string
		printer.Println("AGREE[Y] OR DISAGREE[N]", color.FgMagenta)
		fmt.Scanln(&answer)
		clearSpace()
		err = checkAnswer(answer, metadata, fileUtil)
		if err != nil {
			return err
		}
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

func setUpFile(fileUtil file.File) error {
	_, err := fileUtil.ReadDirectory(constants.ResultDirectory)
	if err != nil {
		err = fileUtil.CreateDirIfNotExist(constants.ResultDirectory)
		if err != nil {
			return err
		}
	}
	exists, err := fileUtil.IsFileExist(constants.ResultFilePath)
	if err != nil {
		return err
	}
	if !exists {
		err = fileUtil.CreateFile(constants.ResultFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func getLatestValueFromFile(fileContent string) (int64, error) {
	if fileContent == "" {
		return 0, nil
	}
	resultArr := strings.Split(fileContent, " ")
	latestString := resultArr[len(resultArr)-1]
	latestValue, err := strconv.ParseInt(latestString, 10, 32)
	if err != nil {
		return -1, err
	}
	return latestValue, nil
}

func updateFile(value int64, fileContent string, fileUtil file.File) error {
	fileContent = fileContent + " " + strconv.FormatInt(value, 10)
	err := fileUtil.WriteFile(constants.ResultFilePath, fileContent)
	return err
}

func clearSpace() {
	fmt.Println()
	fmt.Println()
}

func checkAnswer(answer string, metadata *model.Metadata, fileUtil file.File) error {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	fileContent, err := fileUtil.ReadFile(constants.ResultFilePath)
	if err != nil {
		return err
	}
	latestValue, err := getLatestValueFromFile(fileContent)
	if err != nil {
		return err
	}

	if (strings.EqualFold(answer, "Y") && metadata.Answer) || (strings.EqualFold(answer, "N") && !metadata.Answer) {
		options.FontColor = []figlet4go.Color{
			figlet4go.ColorYellow,
			figlet4go.ColorCyan,
			figlet4go.ColorGreen,
		}
		printer.Println(constants.SuccessMessage, color.FgGreen)
		renderStr, err := ascii.RenderOpts(constants.NneNinja, options)
		if err != nil {
			return err
		}
		latestValue = latestValue + 1
		err = updateFile(latestValue, fileContent, fileUtil)
		if err != nil {
			return err
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
			return err
		}
		latestValue = latestValue - 1
		err = updateFile(latestValue, fileContent, fileUtil)
		if err != nil {
			return err
		}
		fmt.Println(renderStr)
	}
	printer.Println("Things to take care:", color.Underline, color.FgHiWhite)
	clearSpace()
	printer.Println(metadata.Rules, color.FgHiYellow)
	return nil
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
