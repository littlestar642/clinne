package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// File interface define contract for file io
type File interface {
	GetIoReader(string) (io.Reader, error)
	IsFileExist(string) (bool, error)
	ReadFile(string) (string, error)
	ReadDirectory(string) ([]string, error)
	CreateDirIfNotExist(string) error
	CreateFile(string) error
	WriteFile(string, string) error
	GetUserInput() (string, error)
	DeleteFile(filePath string) error
}

type fileUtil struct{}

// New create new instance
func New() File {
	return &fileUtil{}
}

// GetUserInput to take user input
func (f *fileUtil) GetUserInput() (string, error) {
	in := bufio.NewReader(os.Stdin)
	inp, err := in.ReadString('\n')
	if err != nil {
		return "", nil
	}
	return inp, nil
}

func (f *fileUtil) GetIoReader(filePath string) (io.Reader, error) {
	ioReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioReader, nil
}

func (f *fileUtil) IsFileExist(filePath string) (bool, error) {
	dirPath, err := os.Getwd()
	if err != nil {
		return false, err
	}
	filePath = dirPath + filePath
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (f *fileUtil) ReadFile(filePath string) (string, error) {
	dirAbsPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = dirAbsPath + "/" + filePath
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (f *fileUtil) ReadDirectory(dirPath string) ([]string, error) {
	dirAbsPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dirPath = dirAbsPath + dirPath
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	var filePath []string
	for _, f := range files {
		filePath = append(filePath, dirPath+"/"+f.Name())
	}

	return filePath, nil
}

func (f *fileUtil) CreateDirIfNotExist(dir string) error {
	dirPath, err := os.Getwd()
	if err != nil {
		return err
	}
	dir = dirPath + "/" + dir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fileUtil) CreateFile(filePath string) error {
	dirPath, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath = dirPath + "/" + filePath
	_, err = os.Create(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (f *fileUtil) WriteFile(filePath string, content string) error {
	dirPath, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath = dirPath + "/" + filePath
	contentInByte := []byte(content)
	err = ioutil.WriteFile(filePath, contentInByte, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (f *fileUtil) DeleteFile(filePath string) error {
	dirPath, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath = dirPath + "/" + filePath
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
