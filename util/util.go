package util

import (
	"bufio"
	"os"
)

func WriteTo(fileName, data string) {
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0644)
	value := data + "\n"
	defer file.Close()

	_, err := file.WriteString(value)
	if err != err {

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

func GetTextLineValue(filename string, lineNumber int) string {
	formattedFilename := filename + ".txt"
	var savedText string

	file, _ := os.Open(formattedFilename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; i < lineNumber && scanner.Scan(); i++ {
		savedText = scanner.Text()
	}
	return savedText
}
