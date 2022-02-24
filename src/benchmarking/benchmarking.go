/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package benchmarking

import (
	"main/hashing"
	"main/cracking"
	"main/time"
)

// To test benchmarking we need to do the same processing that normal cracking does

var globalState cracking.CrackState;

func crackHash(plaintext string) (found string) {
	hash := hashing.Hash(plaintext, globalState.Algorithm);

	for _, hashed := range globalState.Passwords {
		if (hashed == hash) {
			found = plaintext;

			break;
		}
	}

	return found;
}

func handleFound(found string, index int) {
	if (globalState.RemoveFound && len(globalState.Passwords) > index) {
		globalState.Passwords[index] = globalState.Passwords[len(globalState.Passwords) - 1];
		globalState.Passwords = globalState.Passwords[:len(globalState.Passwords) - 1];
	}
}

func Benchmark(state cracking.CrackState) BenchmarkState {
	globalState = state;
	benchmarkState := BenchmarkState{
		StartTime: time.Seconds(),
	};

	for i := 0; i < len(state.Dictionary); i++ {
		if (time.Seconds() - benchmarkState.StartTime >= state.BenchmarkTime) { break }

		plaintext := state.Dictionary[i];
		cracked := crackHash(plaintext);
		
		if (cracked != "") {
			benchmarkState.Hashed++;
		}
	}

	benchmarkState.EndTime = time.Seconds();

	return benchmarkState;
}