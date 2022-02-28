/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package loading

import (
	"net/http"
	"main/io"
	"fmt"
	"strings"
	"os"
	IO "io"
)

func loadFromHttp(url string) []string {
	response, err := http.Get(url);

	if (err != nil) {
		fmt.Println(err);
		os.Exit(1);
	}

	defer response.Body.Close();

	body, err := IO.ReadAll(response.Body);

	if (err != nil) {
		fmt.Println(err);
		os.Exit(1);
	}

	return strings.Split(string(body), "\n");
}

func Load(file string) []string {
	if (strings.Contains(file, "http")) {
		return loadFromHttp(file);
	}

	return io.ReadFileLines(file);
}