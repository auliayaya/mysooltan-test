package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/auliayaya/mysooltan/internal/model"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func convert2JSON(fileLines []string) (string, error) {
	data := make([]model.JSONModel, 0)
	for _, line := range fileLines {
		split := strings.Split(line, " ")
		data = append(data, model.JSONModel{Column: split[0], Column1: split[1], Column2: split[2], Column3: split[3]})
	}
	bts, err := json.Marshal(data)
	if err != nil {
		check(err)
		return "", err
	}
	js, err := jsonPrettier(string(bts))
	if err != nil {
		check(err)
		return "", err
	}

	return string(js), nil
}

func convert2Text(fileLines []string) (string, error) {

	return strings.Join(fileLines, "\n"), nil
}
func jsonPrettier(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		check(err)
		return "", err
	}
	return prettyJSON.String(), nil
}

func randomString(n int, typ string) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("%s.%s", string(s), typ)
}
func getCurDir() string {
	mydir, err := os.Getwd()
	if err != nil {
		check(err)
	}
	fmt.Println(mydir)
	return mydir
}

func createDir(path string) error {
	_, err := os.Stat(path)

	if err != nil {
		check(err)

		return err
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			check(err)

			return err
		}

	}
	return nil
}
