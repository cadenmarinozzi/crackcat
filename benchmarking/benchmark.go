package benchmarking

import (
	"main/hashing"
	"main/time"
)


func Benchmark(maxTime int, hashingAlgorithm string, nThreads int, startTime int) (int, int) {
	nHashed := 0;

	/*
	* Start nThreads goroutines and split up the workload for the CPU
	**/
	for i := 0; i < nThreads; i++ {
		go func() {
			for (time.Seconds() - startTime / 1000 <= maxTime) {
				hashing.Hash("password123", hashingAlgorithm); // Perform a hash for delay
				nHashed++;
			}
		}();
	}

	for (time.Seconds() - startTime / 1000 <= maxTime) {} // Since goroutines aren't executed on this thread, it won't delay the return

	return nHashed, time.Milliseconds();
}