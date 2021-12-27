package cracking

import (
    "main/time"
    "main/hashing"
    "fmt"
    "strings"
)

func DictionaryFrontBack(dictionary []string, passwords []string, startTime int, maxTime int, logFound bool, removeFound bool, hashingAlgorithm string, usernameSplitter string, nThreads int) (found []string, iterations int, endTime int) {
    nDictionary := len(dictionary) - 1;
    deltaIndex := int(nDictionary / nThreads);
    nPasswords := len(passwords);
    index := 0;

    /*
	* Start nThreads goroutines and split up the workload for the CPU
	**/
    for i := 0; i < nThreads; i++ {
        go func(j int) {
            /*
            * The math behind this:
            * Let nThreads equal 3, meaning we need to split the dictionary into 3 parts and tell each thread where to start and end
            * First we calculate the number of entries each thread needs to handle by taking the number of entries and dividing it by the number of threads
            * Then we get the start index by multiplying the current thread index by the deltaIndex (Number of entries each thread should handle). -
            so for the first thread it's 0 * 2 (0) and for the second thread it'S 1 * 2 (2)
            * and the end index by taking the current thread index, multiplying it by the deltaIndex and adding then adding the deltaIndex -
            so for the first thread it's 0 * 2 + 2 and for the second thread it's 1 * 2 + 2 (4)
            * 
            * j * 2 adds 2 each time so:  | 1 | 2 | 3 |
            *                             v   v   v   v
            *                             0 1 2 3 4 5 6
            *                             a|b|c|d|e|f|g
            *
            * j * 2 + 2 adds 2 each time     ^   ^   ^
            and starts with 2 so:          1 | 2 | 3 |
            **/
            for k := j * deltaIndex; j < j * deltaIndex + deltaIndex; k++ {
                index = k;

                if (time.Seconds() - startTime / 1000 >= maxTime || k > nDictionary) { break } // Stop cracking if the time passed since starting is greater than the max time to crack for

                password := dictionary[k];
                hash := hashing.Hash(password, hashingAlgorithm);
                iterations++;

                for m, hashed := range passwords {
                    if (usernameSplitter == "") {
                        if (hashed == hash) {
                            found = append(found, password);

                            if (removeFound) {
                                passwords[m] = passwords[len(passwords) - 1];
                                passwords = passwords[:len(passwords) - 1];
                            }

                            if (logFound) {
                                fmt.Println(password);
                            }
                        }

                        continue;
                    }
                    
                    details := strings.Split(hashed, usernameSplitter);
                    username := details[0];

                    if (details[1] == hash) {
                        newDetails := username + usernameSplitter + password;
                        found = append(found, newDetails);

                        if (removeFound) {
                            passwords[m] = passwords[len(passwords) - 1];
                            passwords = passwords[:len(passwords) - 1];
                        }

                        if (logFound) {
                            fmt.Println(newDetails);
                        }
                    }
                }
            }
        }(i); // Since all of the goroutines are executed on different threads, we pass the thread index of when the thread is executed instead of the current thread index
    }

    for (time.Seconds() - startTime / 1000 < maxTime && len(found) <= nPasswords && index <= nDictionary) { } // Since goroutines aren't executed on this thread, it won't delay the return

    return found, iterations, time.Milliseconds();
}