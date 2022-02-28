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
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func loadFromHttp(url string) []string {
	response, err := http.Get(url);

	if (response.StatusCode == 200) {
		if (err != nil) {
			fmt.Println(err);
			os.Exit(1);
		}

		defer response.Body.Close();

		bar := progressbar.NewOptions(int(response.ContentLength), progressbar.OptionSetWriter(ansi.NewAnsiStdout()), progressbar.OptionShowBytes(true),  progressbar.OptionSetWidth(15),
			progressbar.OptionSetDescription("Downloading files..."), progressbar.OptionSetTheme(progressbar.Theme{ Saucer: "█", SaucerPadding: " ", BarStart: "|", BarEnd: "|" }));
		data := IO.TeeReader(response.Body, bar);
		
		body, err := IO.ReadAll(data);

		if (err != nil) {
			fmt.Println(err);
			os.Exit(1);
		}

		fmt.Println("\n");

		return strings.Split(string(body), "\n");
	}

	fmt.Println("Failed to download 1 or more files. Check to make sure that the url is correct");
	os.Exit(1);

	return []string{};
}

func Load(file string) []string {
	if (strings.Contains(file, "http")) {
		return loadFromHttp(file);
	}

	return io.ReadFileLines(file);
}