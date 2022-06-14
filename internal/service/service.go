package service

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/auliayaya/mysooltan/internal"
)

type Service struct{}

func NewService() internal.Service {
	return &Service{}
}

// this function will get file error log
func (_s *Service) GetLogFile(path string) ([]string, error) {
	// Get Log File Process from path and return

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	// Reading from a file using reader.
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)
	var fileLines []string
	for reader.Scan() {
		fileLines = append(fileLines, reader.Text())
	}
	file.Close()
	// fmt.Println(fileLines)

	return fileLines, nil
}

// this function will convert error log file by type
func (_s *Service) ConvertLogFile(path, typ, saveto string) (string, error) {
	// Call GetLogFile Service
	fileLines, err := _s.GetLogFile(path)
	check(err)

	switch typ {
	case "json":
		fmt.Println("Convert to JSON")
		// fmt.Println(fileLines)
		res, err := convert2JSON(fileLines)
		if err != nil {
			check(err)
		}
		// fmt.Println(res)
		res2, err := _s.SaveLog(saveto, res, randomString(6, "json"))
		if err != nil {
			check(err)
		}
		fmt.Println(res2)
		return res2, nil
	case "text":
		fmt.Println("Convert to TEXT")
		res, err := convert2Text(fileLines)
		if err != nil {
			check(err)
		}
		// fmt.Println(res)
		res2, err := _s.SaveLog(saveto, res, randomString(6, "txt"))
		if err != nil {
			check(err)
		}
		fmt.Println(res2)
		return res2, nil
	default:
		fmt.Println("Convert to TEXT")
		res, err := convert2Text(fileLines)
		if err != nil {
			check(err)
		}
		// fmt.Println(res)
		res2, err := _s.SaveLog(saveto, res, randomString(6, "txt"))
		if err != nil {
			check(err)
		}
		fmt.Println(res2)
		return res2, nil
	}

}

var (
	stdOufFile string
)

// this function will save log file to the destination location
func (_s *Service) SaveLog(path string, data string, filename string) (string, error) {
	// fmt.Println("Save Log File")

	filePath, err := filepath.Abs(path)

	if err != nil {
		check(err)
	}

	err = createDir(filePath)
	if err != nil {
		check(err)
	}
	stdOufFile = filePath + "/" + filename

	file, err := os.Create(stdOufFile)
	if err != nil {
		check(err)
		return "", err
	}
	_, err = fmt.Fprintln(file, data)
	if err != nil {
		check(err)
		return "", err
	}

	msg := fmt.Sprintf("File succesfully saved to %s", stdOufFile)

	return msg, nil
}
