package main

import (
	"main/time"
    "main/cracking"
	"main/hashing"
    "main/optimization"
	"main/benchmarking"
	genRules "main/rules"
	"main/output"
	"main/errs"
	"main/units"
	"io/ioutil"
	"flag"
	"strings"
	"log"
	fTime "time"
	//"strconv"
	"fmt"
	"os"
	"github.com/shirou/gopsutil/cpu"
)

const WATERMARK = "                                        /$$                             /$$    \n                                       | $$                            | $$    \n  /$$$$$$$  /$$$$$$  /$$$$$$   /$$$$$$$| $$   /$$  /$$$$$$$  /$$$$$$  /$$$$$$  \n /$$_____/ /$$__  $$|____  $$ /$$_____/| $$  /$$/ /$$_____/ |____  $$|_  $$_/  \n| $$      | $$  \\__/ /$$$$$$$| $$      | $$$$$$/ | $$        /$$$$$$$  | $$    \n| $$      | $$      /$$__  $$| $$      | $$_  $$ | $$       /$$__  $$  | $$ /$$\n|  $$$$$$$| $$     |  $$$$$$$|  $$$$$$$| $$ \\  $$|  $$$$$$$|  $$$$$$$  |  $$$$/\n \\_______/|__/      \\_______/ \\_______/|__/  \\__/ \\_______/ \\_______/   \\___/\n";
const KITTY = "       _                        \n       \\`*-.                    \n        )  _`-.                 \n       .  : `. .                \n       : _   '  \\               \n       ; *` _.   `*-._          \n       `-.-'          `-.       \n         ;       `       `.     \n         :.       .        \\    \n         . \\  .   :   .-'   .   \n         '  `+.;  ;  '      :   \n         :  '  |    ;       ;-. \n         ; '   : :`-:     _.`* ;\n      .*' /  .*' ; .*`- +'  `*' \n     `*-*   `*-*  `*-*'";

func main() {
	fmt.Println("crackcat (v1.0.0) starting...\n");

	CPU, _ := cpu.Info();
	cpuInfo := CPU[0];
	cores := int(cpuInfo.Cores);
	sessionName := flag.String("session_name", "crackcat", "");
	rulesFile := flag.String("rules_file", "example/rules.txt", "");
	dictionaryOutputFile := flag.String("dictionary_output_file", "dictionary.txt", "")
	passwordsFile := flag.String("passwords_file", "example/password_list.txt", "");
	dictionaryFile := flag.String("dictionary_file", "example/dictionary.txt", "");
	outputFile := flag.String("output_file", "found.json", "");
	crackType := flag.String("crack_type", "dictionary", "");
	crackMode := flag.String("crack_mode", "front-back", "");
    optimizeDictionary := flag.Bool("optimize_dictionary", false, "");
	maxTime := flag.Int("max_time", 120, "");
	logFound := flag.Bool("log_found", false, "");
	removeFound := flag.Bool("remove_found", true, "");
	hashingAlgorithm := flag.String("hashing_algorithm", "sha256", "");
	hashPasswords := flag.Bool("hash_passwords", false, "");
	usernameSplitter := flag.String("username_splitter", "", "");
    benchmark := flag.Bool("benchmark", false, "");
	watermark := flag.Bool("watermark", false, "");
	nThreads := flag.Int("n_threads", cores, "");
	flag.Parse();

	if (*watermark) {
		fmt.Println(WATERMARK);
		fmt.Println(KITTY);
		fmt.Println("\t\t\tby: nekumelon");
	}

	logFile, err := os.OpenFile("crackcat.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666);

    if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			log.Println(pathError);
		}

		return;
	}

    log.SetOutput(logFile);

	if (*nThreads > cores) {
		fmt.Println("The number of threads is greater than the number of CPU cores. For future sessions, make sure this does not happen");
	}

	fmt.Printf("* CPU: %s, %d MHz, %d cores\n", strings.Split(cpuInfo.ModelName, "           ")[0], int(cpuInfo.Mhz), cores);
	fmt.Printf("crackcat session \"%s\" started\n\n", *sessionName);

	if (*benchmark) {	
		startTime := time.Milliseconds();
		nHashed, endTime := benchmarking.Benchmark(*maxTime, *hashingAlgorithm, *nThreads, startTime);
		hashRateString := units.HashRate(nHashed / 1000000 / *maxTime);
		deltaTimeString := units.LongHandTime(endTime - startTime);
		fmt.Printf("Finished benchmarking in %s. Hash rate: %d\n", deltaTimeString, hashRateString);

		return;
	}

	body, err := ioutil.ReadFile(*rulesFile);

	if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			log.Println(pathError);
		}

		return;
	}
	
	bodyString := string(body);
	splitter := "\n";

	if (strings.Contains(bodyString, "\r\n")) {
		splitter = "\r\n";
	}

	rules := strings.Split(bodyString, splitter);
	body, err = ioutil.ReadFile(*passwordsFile);

	if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			log.Println(pathError);
		}

		return;
	}
	
	bodyString = string(body);
	splitter = "\n";

	if (strings.Contains(bodyString, "\r\n")) {
		splitter = "\r\n";
	}

	passwords := strings.Split(bodyString, splitter);
	body, err = ioutil.ReadFile(*dictionaryFile);

	if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			log.Println(pathError);
		}

		return;
	}

	bodyString = string(body);
	splitter = "\n";

	if (strings.Contains(bodyString, "\r\n")) {
		splitter = "\r\n";
	}

	dictionary := strings.Split(bodyString, splitter);

	if (*crackType == "rules") {
		dictionary = genRules.GenDictionary(dictionary, rules);
	}

	if (*optimizeDictionary) {
        if (*crackMode =="front-back") {
            dictionary = optimization.OptimizeFrontBack(dictionary);
        }
    }

	passwordsSize := 0;
	dictionarySize := 0;

	for _, password := range passwords {
		passwordsSize += len(password);
	}

	for _, password := range dictionary {
		dictionarySize += len(password);
	}

	fmt.Printf("Rules: %d entries\n", len(rules));
	fmt.Printf("Passwords: %d bytes, %d entries\n", passwordsSize, len(passwords));
	fmt.Printf("Dictionary: %d bytes, %d entries\n\n", dictionarySize, len(dictionary));

	if (*crackType == "rules" || *optimizeDictionary) {
		output.Save(*dictionaryOutputFile, *sessionName, dictionary);
	}

	if (*hashingAlgorithm == "" || *hashingAlgorithm == "auto-detect") {
		averageLength := 0;

		for _, password := range passwords {
			averageLength += len(password);
		}

		averageLength /= len(passwords);

		for algorithm, length := range hashing.Lengths {
			if (length == averageLength) {
				*hashingAlgorithm = algorithm;

				break;
			}
		}

		if (*hashingAlgorithm == "" || *hashingAlgorithm == "auto-detect") {
			return;
		}
	}

	if (*hashPasswords) {
		nPasswords := len(passwords);

		if (nPasswords > 50) {
			averageLength := 0;

			for i := 0; i < 50; i++ {
				averageLength += len(passwords[i]);
			}

			averageLength /= 50;

			if (averageLength == hashing.Lengths[*hashingAlgorithm]) {
				fmt.Println("It looks like the password list is already hashed. Run crackcat without hash passwords argument to prevent this");
			}
		}

		for i := 0; i < nPasswords; i++ {
			if (*usernameSplitter == "") {
				passwords[i] = hashing.Hash(passwords[i], *hashingAlgorithm);
			} else {
				details := strings.Split(passwords[i], *usernameSplitter);
				passwords[i] = details[0] + *usernameSplitter + hashing.Hash(details[1], *hashingAlgorithm);
			}
		}
	}

    var found []string;
    iterations := 0;
    endTime := 0;
	startTime := time.Milliseconds();
	fStartTime := fTime.Now();
	nPasswords := len(passwords);
	nDict := len(dictionary) - 1;

	if (*nThreads > nDict) {
		*nThreads = nDict;
	}

	startTime = time.Milliseconds();
	nHashed := 0;

	for (time.Seconds()  - startTime / 1000 <= 1) {
		hashing.Hash("password123", *hashingAlgorithm);
		nHashed++;
	}

	fmt.Printf("Estimated time..: %s\n", units.ShortHandTime(int((float64(len(dictionary)) / float64(nHashed)) * 1000)));
	
	switch (*crackType) {
		case ("dictionary"):
			switch (*crackMode) {
				case ("front-back"):
					found, iterations, endTime = cracking.DictionaryFrontBack(dictionary, passwords, startTime, *maxTime, *logFound, *removeFound, *hashingAlgorithm, *usernameSplitter, *nThreads);
			}

		case ("rules"):
			switch (*crackMode) {
				case ("front-back"):
					found, iterations, endTime = cracking.DictionaryFrontBack(dictionary, passwords, startTime, *maxTime, *logFound, *removeFound, *hashingAlgorithm, *usernameSplitter, *nThreads);
			}
	}

	fEndTime := fTime.Now();
	deltaTimeString := units.ShortHandTime(endTime - startTime);
	nFound := len(found);
	foundPercent := int((float64(nFound) / float64(nPasswords)) * 100);
	output.Save(*outputFile, *sessionName, found);

	fmt.Printf("Start time......: %s\n", fStartTime.Format("03:04:05"));
	fmt.Printf("End time........: %s\n", fEndTime.Format("03:04:05"));
	fmt.Printf("Time taken......: %s\n", deltaTimeString);
	fmt.Printf("Iterations......: %d\n", iterations);
	fmt.Printf("Found...........: %d passwords (%d%%)\n", nFound, foundPercent);
	fmt.Printf("Algorithm.......: %s", *hashingAlgorithm);
}