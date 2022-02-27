/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package main

import (
	"main/terminal"
	"main/io"
	"main/cracking"
	"main/benchmarking"
	"main/hashing"
	"main/time"
	genRules "main/rules"
	"main/optimization"
	"flag"
	"github.com/shirou/gopsutil/cpu"
	"fmt"
	Ftime "time"
	"strings"
	"os"
)

var version string = "1.2.0";

func main() {
	CPU, _ := cpu.Info();
	cores := int(CPU[0].Cores);

	// Handle all of the flags and terminal args
	threads := flag.Int("threads", cores, "The number of threads to distribute the workload on");

	passwordsFile := flag.String("passwords", "./examples/big_passwords1.txt", "The file holding the passwords to crack");
	sessionName := flag.String("session", "crackcat", "The name of the crackcat session");
	dictionaryFile := flag.String("dictionary", "./examples/big_dictionary1.txt", "The file holding the passwords to check against");
	rulesFile := flag.String("rules", "", "The file holding the rules to modify the dictionary with");

	dictionaryCutoff := flag.Int("cutoff", 0, "The limit to how many dictionary entries are used");

	algorithm := flag.String("algorithm", "sha256", "The hashing algorithm to hash with");

	prehash := flag.Bool("prehash", false, "Whether to pre-hash the passwords for a plaintext list");

	watermark := flag.Bool("watermark", true, "Whether to show the watermark in the terminal");
	kitty := flag.Bool("kitty", false, "Whether to show the kitty");

	device := flag.String("device", "CPU", "The device to use for cracking (CPU or GPU)")

	optimizeEntries := flag.Bool("optimize_entries", true, "Whether to optimize the dictionary/rules entries");

	maxTime := flag.Int("max_time", 60, "The maximum time to crack for");
	benchmarkTime := flag.Int("benchmark_time", 1, "The maximum time to benchmark for");

	crackingMethod := flag.String("cracking_method", "dictionary", "The method to use to crack the hashes");
	crackingMode := flag.String("cracking_mode", "left-right", "The direction mode to use to crack the hashes");

	logFound := flag.Bool("log_found", false, "Whether to log the found passwords to the terminal");
	sameLineLogs := flag.Bool("sameline_logs", true, "Whether logged passwords should overwrite the last logged password line");
	removeFound := flag.Bool("remove_found", true, "Whether to remove the found passwords from the list to increase performance"); // I don't really know why you wouldn't want to do this
	
	flag.Parse();

	state := cracking.CrackState{
		Algorithm: *algorithm,
		MaxTime: *maxTime,
		Threads: *threads,
		CrackingMethod: *crackingMethod,
		CrackingMode: *crackingMode,
		LogFound: *logFound,
		RemoveFound: *removeFound,
		SessionName: *sessionName,
		SameLineLogs: *sameLineLogs,
		BenchmarkTime: *benchmarkTime,
		FormattedStartTime: Ftime.Now().Format("03:04:05 PM"),
	};

	io.CreateDirectory("./" + *sessionName);
	
	terminal.Header(version);

	if (*watermark) {
		terminal.Watermark(*kitty); // kitty :3
	}

	terminal.Devices();

	// Causes a paradox if you need to know the hashing algorithm to prehash but also need to detect the algorithm
	if (*prehash && *algorithm == "auto") {
		fmt.Println("Cannot prehash if the algorithm is not defined");
		os.Exit(1);
	}

	// This can cause some of the goroutines to be ignored so it's considered dangerous
	if (*threads > cores) {
		fmt.Println("warning: The number of threads supplied is greater than the number of cores on the device, this could cause performance issues\n");
	}

	if (*logFound) {
		fmt.Println("warning: Logging found passwords results in crackcat cracking slower, log_found should not be used for general purposes\n");
	}

	terminal.Optimizers(map[string]bool{
		"optimize_entries": *optimizeEntries,
		"remove_found": *removeFound,
	});
	
	// Load the passwords and dictionary entries from the files
	dictionary := io.ReadFileLines(*dictionaryFile);

	if (*dictionaryCutoff != 0) {
		dictionary = dictionary[:*dictionaryCutoff];
	}	

	dictionarySize := len(dictionary);

	passwords := io.ReadFileLines(*passwordsFile);
	passwordsSize := len(passwords);

	if (*optimizeEntries) {
		passwords = optimization.RemoveEmptyHashes(passwords, *algorithm);
		fmt.Printf("Passwords: %d entries, %d optimized entries\n", passwordsSize, len(passwords));
	} else {
		fmt.Printf("Passwords: %d entries\n", passwordsSize);
	}

	if (*rulesFile != "") {
		rules := io.ReadFileLines(*rulesFile);
		rulesSize := len(rules);

		if (*optimizeEntries) {
			rules = optimization.RemoveDuplicates(rules);
			fmt.Printf("Rules: %d rules, %d optimized rules\n", rulesSize, len(rules));

			dictionary = optimization.RemoveDuplicates(genRules.GenerateRules(dictionary, rules));
			fmt.Printf("Dictionary: %d entries, %d optimized entries\n", dictionarySize, len(dictionary));
		} else {
			fmt.Printf("Rules: %d rules\n", rulesSize);
			fmt.Printf("Dictionary: %d entries", dictionarySize);
		}

		io.WriteFile("./" + *sessionName + "/dictionary_" + Ftime.Now().Format("01-02-2006 03_04_05") + ".txt", strings.Join(dictionary, "\n"));
	} else if (*optimizeEntries) {
		dictionary = optimization.RemoveDuplicates(dictionary);
		fmt.Printf("Dictionary: %d entries, %d optimized entries\n", dictionarySize, len(dictionary));
		io.WriteFile("./" + *sessionName + "/dictionary_" + Ftime.Now().Format("01-02-2006 03_04_05") + ".txt", strings.Join(dictionary, "\n"));
	} else {
		fmt.Printf("Dictionary: %d entries", dictionarySize);
	}

	fmt.Println();

	detectedAlgorithm := hashing.DetectAlgorithm(passwords[0]);

	if (*algorithm == "auto") {
		*algorithm = detectedAlgorithm;
		state.Algorithm = *algorithm;
		fmt.Printf("Auto detected %s as the hashing algorithm\n\n", detectedAlgorithm);
	} else if (detectedAlgorithm != *algorithm && !*prehash) {
		fmt.Println("warning: The specified hashing algorithm does not match the detected hashing algorithm. Make sure the supplied algorithm matches the specified algorithm\n");
	}

	if (*prehash) {
		for i := 0; i < len(passwords); i++ {
			passwords[i] = hashing.Hash(passwords[i], *algorithm);
		}

		io.WriteFile("./" + *sessionName + "/passwords_" + Ftime.Now().Format("01-02-2006 03_04_05") + ".txt", strings.Join(passwords, "\n"));
	}

	if (*threads > len(dictionary)) {
		fmt.Println("warning: The number of threads is larger than the number size of the dictionary so it has been normalized\n");
		*threads = len(dictionary);
	}

	// Start cracking

	state.Threads = *threads;
	state.Passwords = passwords;
	state.NPasswords = len(state.Passwords);
	state.Dictionary = dictionary;

	benchmarkState := benchmarking.Benchmark(state);
	state.EstimatedTime = len(state.Passwords) / benchmarkState.Hashed / state.BenchmarkTime;

	fmt.Printf("Started cracking at %s\n\n", state.FormattedStartTime);

	fmt.Printf("Estimated time..: %d seconds\n", state.EstimatedTime);

	if (*device == "CPU") {
		state = cracking.Crack(state);

		state.EndTime = time.Seconds();
		state.FormattedEndTime = Ftime.Now().Format("03:04:05 PM");
		
		io.WriteFile("./" + state.SessionName + "/found_" + Ftime.Now().Format("01-02-2006 03_04_05") + ".txt", strings.Join(state.Found, "\n"));
		terminal.Cracked(state);
	} else {
		fmt.Println("crackcat doesn't support this device for cracking yet");
		os.Exit(1);
	}
}