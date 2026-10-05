package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetInput() (input string) {
	reader := bufio.NewReader(os.Stdin)
	unformattedInput, _ := reader.ReadString('\n')
	input = strings.Trim(unformattedInput, "\r\n")
	return
}

func CreateFile(file string) error {
	dir := filepath.Dir(file) //.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	} else {
		return nil
	}
}

func WriteTo(filePath, data string) {

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	} else {

	}
	defer file.Close()

	value := data + "\n"
	_, writeErr := file.WriteString(value)
	if writeErr != nil {
		fmt.Println("Can't write to this file!")
	}
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}

func FloatToString(dec float64) string {
	converter := fmt.Sprintf("%.2f", dec)
	return converter
}

func StringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(strings.Trim(s, "\r\n"), 64)
	if err != nil {
		return 0.0, errors.New("can't convert to number! ")
	} else {
		return f, nil
	}
}

func StringToInt(s string) (int64, error) {
	stripString := strings.Trim(s, "\r\n")
	num, err := strconv.ParseInt(stripString, 10, 64)
	if err != nil {
		return 0, errors.New("can't convert input to integer")
	} else {
		return num, nil
	}
}
func IntToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

func GetTextLineValue(filePath string, lineNumber int) (savedText string) {
	formattedFilename := filePath

	file, _ := os.Open(formattedFilename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	realValue := ""
	selectedValue := ""

	if lineNumber < 0 {
		for scanner.Scan() {
			selectedValue = scanner.Text()
			if selectedValue == "\n" || selectedValue == "\r\n" || selectedValue == "" {
				return realValue
			} else {
				realValue = selectedValue
			}
		}
		return realValue
	} else {
		for i := 0; i < lineNumber && scanner.Scan(); i++ {
			selectedValue = scanner.Text()
			if selectedValue == "\n" || selectedValue == "\r\n" || selectedValue == "" {
				return realValue
			} else {
				realValue = selectedValue
			}
		}
		return realValue
	}

}
