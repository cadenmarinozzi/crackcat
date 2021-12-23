package benchmarking

import (
	"main/hashing"
	"main/time"
)

/*
* Benchmark starts n threads that create hashes until the maxTime has been reached. nHashed is the amount of hashes that got hashed in those n seconds
*/
func Benchmark(maxTime int, hashingAlgorithm string, nThreads int, startTime int) (int, int) {
	nHashed := 0;

	for i := 0; i < nThreads; i++ {
		go func() {
			for (time.Seconds()  - startTime / 1000 <= maxTime) {
				hashing.Hash("password123", hashingAlgorithm);
				nHashed++;
			}
		}()
	}

	for (time.Seconds() - startTime / 1000 <= maxTime) {} // This is here so that the function doesn't return until the maxTime is reached

	return nHashed, time.Milliseconds();
}