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
)

var globalState CrackState; // No need to keep passing the state around in the same file

func crackHash(plaintext string) (found string, index int) {
	hash := hashing.Hash(plaintext, globalState.Algorithm);

	for i, hashed := range globalState.Passwords {
		if (hashed == hash) {
			found = plaintext;
			index = i;

			break;
		}
	}

	return found, index;
}

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

func process() {
	globalState.StartTime = time.Seconds();
	deltaIndex := int(math.Ceil(float64(len(globalState.Dictionary) - 1)) / float64(globalState.Threads));

	var threads []*Thread;
	running := true;

	if (globalState.Threads > 1) {
		for i := 0; i < globalState.Threads; i++ {
			endIndex := i * deltaIndex + deltaIndex;

			if (i == globalState.Threads - 1) {
				endIndex = len(globalState.Dictionary);
			}

			padding := 1;

			if (i == 0) {
				padding = 0;
			}

			thread := Thread{
				Index: i,
				EndIndex: endIndex,
				StartIndex: i * deltaIndex + padding,
				Running: true,
			};

			threads = append(threads, &thread);

			go func() {
				for i := thread.StartIndex; i < thread.EndIndex; i++ { // Dictionary entries
					globalState.Iterations++;
					thread.EntryIndex = i;
	
					if (time.Seconds() - globalState.StartTime >= globalState.MaxTime || thread.EntryIndex >= thread.EndIndex - 1) { 
						thread.Running = false; 
					}
	
					var plaintext string;

					switch (globalState.CrackingMode) {
						case ("left-right"):
							plaintext = globalState.Dictionary[i];

							break;

						case ("right-left"):
							plaintext = globalState.Dictionary[len(globalState.Dictionary) - 1 - i];

							break;
					}

					cracked, index := crackHash(plaintext);
					
					if (cracked != "") {
						handleFound(cracked, index);
					}
				}
			}();
		}
	} else {
		thread := Thread{
			Index: 0,
			EndIndex: len(globalState.Dictionary),
			Running: true,
		};

		threads = append(threads, &thread);

		for i := 0; i < thread.EndIndex; i++ { // Dictionary entries
			globalState.Iterations++;
			thread.EntryIndex = i;

			if (time.Seconds() - globalState.StartTime >= globalState.MaxTime || thread.EntryIndex >= thread.EndIndex - 1) { 
				thread.Running = false; 
			}

			plaintext := globalState.Dictionary[i];
			cracked, index := crackHash(plaintext);
			
			if (cracked != "") {
				handleFound(cracked, index);
			}
		}
	}

	for (running) {
		running = false;

		for _, thread := range threads {
			if (thread.Running) {
				running = true;
			}
		}
	}
}


func DictionaryAttack(state CrackState) CrackState {
	globalState = state;
	process();

	return globalState;
}