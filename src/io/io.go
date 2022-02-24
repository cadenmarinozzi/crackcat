/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package io

import (
	"io/ioutil"
	"os"
	"strings"
	"errors"
	"fmt"
)

// I HATE io stuff so all of it is handled here in simple functions

func HandleFileError(err error) {
	if (err == nil) { return }

	fmt.Println(err.Error());
	os.Exit(1);
}

func WriteFile(filePath string, body string) {
	file, err := os.Create(filePath);
	HandleFileError(err);
	file.WriteString(body);
	file.Close();
}

func ReadFile(filePath string) string {
	body, err := ioutil.ReadFile(filePath);
	HandleFileError(err);

	return string(body);
}

func ReadFileLines(filePath string) []string {
	body := ReadFile(filePath);
	seperator := "\n";

	if (strings.Contains(body, "\r\n")) {
		seperator = "\r\n";
	}

	split := strings.Split(body, seperator);

	return split;
}

func CreateDirectory(directoryPath string) {
	if _, err := os.Stat(directoryPath); (errors.Is(err, os.ErrNotExist)) {
        os.Mkdir(directoryPath, 0755);
    }
}