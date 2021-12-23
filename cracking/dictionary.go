package cracking

import (
    "main/time"
    "main/hashing"
    "fmt"
    "strings"
)

func DictionaryFrontBack(dictionary []string, passwords []string, startTime int, maxTime int, logFound bool, removeFound bool, hashingAlgorithm string, usernameSplitter string, nThreads int) ([]string, int, int) {
    var found []string;
    iterations := 0;

    nDict := len(dictionary) - 1;
    K := int(nDict / nThreads);
    nPasswords := len(passwords);
    p := 0;

    for i := 0; i < nThreads; i++ {
        go func(j int) {
            for k := j * K; j < j * K + K; k++ {
                p = k;

                if (time.Seconds() - 1 - startTime / 1000 >= maxTime || k > nDict) {
                    break;
                }

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
                    } else {
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
            }
        }(i)
    }

    for (time.Seconds() - 1 - startTime / 1000 < maxTime && len(found) <= nPasswords && p <= nDict) { }

    return found, iterations, time.Milliseconds();
}