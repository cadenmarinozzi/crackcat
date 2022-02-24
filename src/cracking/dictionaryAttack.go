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
func handleFound(found string, index int) {
	globalState.Found = append(globalState.Found, found);

	if (globalState.LogFound) { // This really shouldn't be multi threaded
		if (globalState.SameLineLogs && len(globalState.Found) > 1) {
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

	/* This really sucks. I need to figure out a better way to do this since it causes some entries-
		to get missed, depending on the thread speed
	*/
	deltaIndex := (len(globalState.Dictionary) + globalState.Threads - 1) / globalState.Threads;
	
	var threads []*Thread;
	running := true;

	if (globalState.Threads > 1) {
		for i := 0; i < globalState.Threads; i++ {
			startIndex := i * deltaIndex;
			endIndex := int(math.Min(float64(startIndex + deltaIndex), float64(len(globalState.Dictionary))));

			thread := Thread{
				Index: i,
				EndIndex: endIndex,
				StartIndex: startIndex,
				Running: true,
			};

			if (i != 0) {
				thread.StartIndex++;
			}

			threads = append(threads, &thread); // We insert the address so that the main thread can access the threads

			go func() { // goroutines oh goroutines I love you very much
				for i := thread.StartIndex; i < thread.EndIndex; i++ {
					var plaintext string;
					thread.EntryIndex = i;

					if (time.Seconds() - globalState.StartTime >= globalState.MaxTime) {
						thread.Running = false;
			
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
						handleFound(cracked, index);
					}

					globalState.Iterations++;
				}

				thread.Running = false;
			}();
		}
	} else {
		thread := Thread{
			Index: 0,
			EndIndex: len(globalState.Dictionary),
			StartIndex: 0,
			Running: true,
		};

		threads = append(threads, &thread); // We insert the address so that the main thread can access the threads

		for i := thread.StartIndex; i < thread.EndIndex; i++ {
			var plaintext string;
			thread.EntryIndex = i;

			if (time.Seconds() - globalState.StartTime >= globalState.MaxTime) {
				thread.Running = false;
	
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
				handleFound(cracked, index);
			}

			globalState.Iterations++;
		}

		thread.Running = false;
	}
	
	for (running) {
		running = false;

		for _, thread := range threads {
			if (thread.Running) {
				running = true;
			}
		}
	}

	return globalState;
}