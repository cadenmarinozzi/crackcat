/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package loading

import (
	"net/http"
	"io"
)

func loadFromHttp(url string) []string {
	response, err := http.Get(url);

	if (err != nil) {
		fmt.Println(err);
		os.Exit(1);
	}

	return strings.Split(response.Body, "\n");
}

func Load(file string) []string {
	if (strings.Contains(file, "http")) {
		return loadFromHttp(file);
	}

	return io.ReadFileLines(file);
}