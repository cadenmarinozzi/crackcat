package output

import (
	"main/errs"
    "os"
    "strings"
    "encoding/json"
    "time"
	"errors"
	"fmt"
)

/*
* Format and save found to the file in the specified directory
*/
func Save(fileName string, directory string, found []string) {
	// If the directory doesn't exist, make it
    if _, err := os.Stat(directory); (errors.Is(err, os.ErrNotExist)) {
        os.Mkdir(directory, 0755);
    }

    fileDetails := strings.Split(fileName, ".");
	extension := "";

	if (len(fileDetails) > 1) {
		extension = fileDetails[1];
	}

    file, err := os.Create(directory + "/" + fileDetails[0] + "_" + time.Now().Format("01-02-2006 03_04_05") + extension); // Create the output file

	if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			fmt.Println(pathError);
		}

		return;
	}

	output := "";

	// Format the output
	if (extension == ".json") {
		foundJson, _ := json.Marshal(found);
		output = string(foundJson);
	} else {
		output = strings.Join(found, "\n");
	}

	file.WriteString(output);
	file.Close();
}