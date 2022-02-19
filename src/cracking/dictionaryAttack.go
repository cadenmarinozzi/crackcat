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

func crackHash(plaintext string, state CrackState) (found string, index int) {
	hash := hashing.Hash(plaintext, state.Algorithm);

	for i, hashed := range state.Passwords {
		if (hashed == hash) {
			found = plaintext;
			index = i;

			break;
		}
	}

	return found, index;
}

func handleFound(found string, index int, state CrackState) CrackState {
	state.Found = append(state.Found, found);

	if (state.LogFound) {
		if (state.SameLineLogs && len(state.Found) > 1) {
			fmt.Printf("\033[1A\033[2K");
		}

		fmt.Printf("%s:%s\n", state.Passwords[index], found);
	}

	if (state.RemoveFound && len(state.Passwords) > index) {
		state.Passwords[index] = state.Passwords[len(state.Passwords) - 1];
		state.Passwords = state.Passwords[:len(state.Passwords) - 1];
	}

	return state;
}

func process(state CrackState) CrackState {
	state.StartTime = time.Seconds();
	state.Iterations = 0;
	deltaIndex := int(math.Ceil(float64(len(state.Dictionary) - 1)) / float64(state.Threads));
	var threads []*Thread;
	running := true;

	if (state.Threads > 1) {
		for i := 0; i < state.Threads; i++ {
			endIndex := i * deltaIndex + deltaIndex;

			if (i == state.Threads - 1) {
				endIndex = len(state.Dictionary);
			}

			thread := Thread{
				Index: i,
				EndIndex: endIndex,
				Running: true,
			};

			threads = append(threads, &thread);

			go func() {
				// This is why we need ternary ops
				padding := 1;

				if (thread.Index == 0) {
					padding = 0;
				}

				for i := thread.Index * deltaIndex + padding; i < endIndex; i++ { // Dictionary entries
					state.Iterations++;
					thread.EntryIndex = i;
	
					if (time.Seconds() - state.StartTime >= state.MaxTime || thread.EntryIndex >= thread.EndIndex - 1) { 
						thread.Running = false; 
					}
	
					var plaintext string;

					switch (state.CrackingMode) {
						case ("left-right"):
							plaintext = state.Dictionary[i];

							break;

						case ("right-left"):
							plaintext = state.Dictionary[len(state.Dictionary) - 1 - i];

							break;
					}

					cracked, index := crackHash(plaintext, state);
					
					if (cracked != "") {
						state = handleFound(cracked, index, state);
					}
				}
			}();
		}
	} else { // I'll handle this later
		thread := Thread{
			Index: 0,
			EndIndex: len(state.Dictionary),
			Running: true,
		};

		threads = append(threads, &thread);

		for i := 0; i < thread.EndIndex; i++ { // Dictionary entries
			state.Iterations++;
			thread.EntryIndex = i;

			if (time.Seconds() - state.StartTime >= state.MaxTime || thread.EntryIndex >= thread.EndIndex - 1) { 
				thread.Running = false; 
			}

			plaintext := state.Dictionary[i];
			cracked, index := crackHash(plaintext, state);
			
			if (cracked != "") {
				state = handleFound(cracked, index, state);
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

	return state;
}


func DictionaryAttack(state CrackState) CrackState {
	state = process(state);

	return state;
}