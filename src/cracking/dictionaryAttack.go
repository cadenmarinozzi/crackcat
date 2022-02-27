/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package cracking

import (
	"main/time"
	"main/hashing"
	"fmt"
	"math"
	"os"
	"sync"
)

var globalState CrackState; // No need to keep passing the state around in the same file

/*
* This is the function that goes through the hashed password list to compare the plaintext-
	entry to the hash
**/
func crackHash(plaintext string) (found string, index int) {
	hash := hashing.Hash(plaintext, globalState.Algorithm);

	for i, hashed := range globalState.Passwords {
		if (hashed == hash) {
			found = plaintext;
			index = i; // We need the index for removing the password later on 

			break;
		}
	}

	return found, index;
}

/*
* Handle the found password. This means logging it, removing it, etc
**/
func handleFound(thread *Thread, found string, index int) {
	thread.Found = append(thread.Found, found);

	if (globalState.LogFound) { // This really shouldn't be multi threaded
		if (globalState.SameLineLogs && len(thread.Found) > 1) {
			fmt.Printf("\033[1A\033[2K");
		}

		fmt.Printf("%s:%s\n", globalState.Passwords[index], found);
	}

	if (globalState.RemoveFound && len(globalState.Passwords) > index) {
		globalState.Passwords[index] = globalState.Passwords[len(globalState.Passwords) - 1];
		globalState.Passwords = globalState.Passwords[:len(globalState.Passwords) - 1];
	}
}

func DictionaryAttack(state CrackState) CrackState {
	globalState = state;
	globalState.StartTime = time.Seconds();

	deltaIndex := (len(globalState.Dictionary) + globalState.Threads - 1) / globalState.Threads;
	
	var threads []*Thread;

	var waitGroup sync.WaitGroup; // Just found out about these lol
	waitGroup.Add(globalState.Threads);

	if (globalState.Threads > 1) {
		for i := 0; i < globalState.Threads; i++ {
			startIndex := i * deltaIndex;
			endIndex := int(math.Min(float64(startIndex + deltaIndex), float64(len(globalState.Dictionary))));

			thread := Thread{
				Index: i,
				EndIndex: endIndex,
				StartIndex: startIndex,
			};

			threads = append(threads, &thread); // We insert the address so that the main thread can access the threads

			go func() {
				defer waitGroup.Done();

				for j := thread.StartIndex; j < thread.EndIndex; j++ {
					var plaintext string;
					thread.EntryIndex = j;

					if (time.Seconds() - globalState.StartTime >= globalState.MaxTime/* || len(globalState.Passwords) == 0*/) {
						break;
					}

					switch (globalState.CrackingMode) {
						case ("left-right"):
							plaintext = globalState.Dictionary[thread.EntryIndex];

							break;

						case ("right-left"):
							plaintext = globalState.Dictionary[len(globalState.Dictionary) - 1 - thread.EntryIndex];

							break;

						default:
							fmt.Println("Unsupported cracking mode");
							os.Exit(1);
					}

					cracked, index := crackHash(plaintext);
					
					if (cracked != "") {
						handleFound(&thread, cracked, index);
					}
				}
			}();
		}
	}

	waitGroup.Wait();

	for _, thread := range threads {
		globalState.Iterations += thread.EntryIndex;
		globalState.Found = append(globalState.Found, thread.Found...);
	}

	return globalState;
}