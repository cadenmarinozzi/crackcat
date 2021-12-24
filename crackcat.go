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
	"strconv"
	"fmt"
	"os"
	"github.com/shirou/gopsutil/cpu"
)

const WATERMARK = "                                        /$$                             /$$    \n                                       | $$                            | $$    \n  /$$$$$$$  /$$$$$$  /$$$$$$   /$$$$$$$| $$   /$$  /$$$$$$$  /$$$$$$  /$$$$$$  \n /$$_____/ /$$__  $$|____  $$ /$$_____/| $$  /$$/ /$$_____/ |____  $$|_  $$_/  \n| $$      | $$  \\__/ /$$$$$$$| $$      | $$$$$$/ | $$        /$$$$$$$  | $$    \n| $$      | $$      /$$__  $$| $$      | $$_  $$ | $$       /$$__  $$  | $$ /$$\n|  $$$$$$$| $$     |  $$$$$$$|  $$$$$$$| $$ \\  $$|  $$$$$$$|  $$$$$$$  |  $$$$/\n \\_______/|__/      \\_______/ \\_______/|__/  \\__/ \\_______/ \\_______/   \\___/\n";
const KITTY = "       _                        \n       \\`*-.                    \n        )  _`-.                 \n       .  : `. .                \n       : _   '  \\               \n       ; *` _.   `*-._          \n       `-.-'          `-.       \n         ;       `       `.     \n         :.       .        \\    \n         . \\  .   :   .-'   .   \n         '  `+.;  ;  '      :   \n         :  '  |    ;       ;-. \n         ; '   : :`-:     _.`* ;\n      .*' /  .*' ; .*`- +'  `*' \n     `*-*   `*-*  `*-*'";

func main() {
	fmt.Println(WATERMARK);
	fmt.Println(KITTY);
	fmt.Println("\t\t\tby: nekumelon\n");
	
	// Get information about the system
	CPU, _ := cpu.Info();
	cpuInfo := CPU[0];
	cores := int(cpuInfo.Cores);

	// The arguments that can be passed to the crackcat command
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
	nThreads := flag.Int("n_threads", cores, "");

	flag.Parse();
	
	// Handle logging for crackcat logs
	logFile, err := os.OpenFile("crackcat.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666);

    if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			fmt.Println(pathError);
		}

		return;
	}

    log.SetOutput(logFile);
	log.Printf("crackcat session '%s' started", *sessionName);

	if (*nThreads > cores) {
		fmt.Println("The number of threads is greater than the number of CPUs. It is recommended to use a number of threads less than or equal to the number of CPUs\n");
	}

	fmt.Printf("CPU: %s\n", cpuInfo.ModelName);
	fmt.Printf("Speed: %s MHz\n", strconv.Itoa(int(cpuInfo.Mhz)));
	fmt.Printf("Cores: %s\n", strconv.Itoa(cores));

	// Handle benchmarks
	if (*benchmark) {
		startTime := time.Milliseconds();
		startTimeString := strconv.Itoa(startTime);
		fmt.Println("Started benchmarking at start time: " + startTimeString);
		log.Println("Started benchmarking at start time: " + startTimeString);

		nHashed, endTime := benchmarking.Benchmark(*maxTime, *hashingAlgorithm, *nThreads, startTime);
		
		hashRateString := units.HashRate(nHashed / 1000000 / *maxTime);
		deltaTimeString := units.LongHandTime(endTime - startTime);
		endTimeString := strconv.Itoa(endTime);
		fmt.Println("Finised benchmarking at " + endTimeString + ". It took " + deltaTimeString + ". Hash rate: " + hashRateString);
		log.Println("Finised benchmarking at " + endTimeString + ". It took " + deltaTimeString + ". Hash rate: " + hashRateString);

		return;
	}

	// Get and parse all of the required files

	body, err := ioutil.ReadFile(*rulesFile);

	if (err != nil) {
		if pathError, ok := err.(*errs.PathError); ok {
			fmt.Println(pathError);

			return;
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
			fmt.Println(pathError);

			return;
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
			fmt.Println(pathError);

			return;
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

	if (*crackType == "rules" || *optimizeDictionary) {
		output.Save(*dictionaryOutputFile, *sessionName, dictionary); // Save the output file
	}

	/**
	* If no hashing algorithm is specified or the hashing algorithm is set to auto detect:
	* Calculate the average length of the password list and check which algorithms output size matches the integer average.
	* The reason the average is taken instead of just taking the first element of the list and getting the size of it is because if there are a few stray passwords that don't match the algorithm, they are ignored.
	*/
	if (*hashingAlgorithm == "" || *hashingAlgorithm == "auto-detect") {
		// Calculate the average (X1 + X2 + X3 + ... Xn / n)
		 
		averageLength := 0;

		for _, password := range passwords {
			averageLength += len(password);
		}

		averageLength /= len(passwords);

		for algorithm, length := range hashing.Lengths {
			if (length == averageLength) {
				*hashingAlgorithm = algorithm;
				fmt.Println("Auto detected " + algorithm + " as the hashing algorithm\n");
				log.Println("Auto detected " + algorithm + " as the hashing algorithm\n");

				break;
			}
		}

		if (*hashingAlgorithm == "" || *hashingAlgorithm == "auto-detect") {
			fmt.Println("Couldn't auto detect a hashing algorithm. Make sure the password list is hashed with a supported hashing algorithm");
			log.Println("Couldn't auto detect a hashing algorithm. Make sure the password list is hashed with a supported hashing algorithm");

			return;
		}
	}

	// If the password list isn't pre hashed (For testing purposes), hash it
	if (*hashPasswords) {
		nPasswords := len(passwords);

		if (nPasswords > 50) {
			averageLength := 0;

			for i := 0; i < 50; i++ {
				averageLength += len(passwords[i]);
			}

			averageLength /= 50;

			if (averageLength == hashing.Lengths[*hashingAlgorithm]) {
				fmt.Println("It looks like the password list is already hashed. In the future, run crackcat without passing the hash_passwords argument\n");
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
	startTimeString := strconv.Itoa(startTime);
	nPasswords := len(passwords);
	nPasswordsString := strconv.Itoa(nPasswords);
	nDict := len(dictionary) - 1;
	nDictString := strconv.Itoa(nDict);

	if (*nThreads > nDict) {
		fmt.Println("The thread count is greater than the amount of entries in the dictionary. Defaulting to a thread count the size of the dictionary, " + nDictString);
		log.Println("The thread count is greater than the amount of entries in the dictionary. Defaulting to a thread count the size of the dictionary, " + nDictString);
		*nThreads = nDict;
	}

	fmt.Println("Started cracking " + nPasswordsString + " passwords at start time: " + startTimeString);
	log.Println("Started cracking " + nPasswordsString + " passwords at start time: " + startTimeString);

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

	deltaTimeString := units.LongHandTime(endTime - startTime);
	endTimeString := strconv.Itoa(endTime);
	nFound := len(found);
	nFoundString := strconv.Itoa(nFound);
	iterationsString := strconv.Itoa(iterations);
	foundPercentString := strconv.Itoa(int((float64(nFound) / float64(nPasswords)) * 100));
	fmt.Println("Finished cracking " + foundPercentString + "% of " + nPasswordsString + " passwords and found " + nFoundString + " at end time: " + endTimeString + ". It took " + deltaTimeString);
	log.Println("Finished cracking " + foundPercentString + "% of " + nPasswordsString + " passwords and found " + nFoundString + " at end time: " + endTimeString + ". It took " + deltaTimeString);
	fmt.Println("Iterations: " + iterationsString);
	log.Println("Iterations: " + iterationsString);
	fmt.Println("Found: " + nFoundString);
	
	output.Save(*outputFile, *sessionName, found); // Save the output file
}