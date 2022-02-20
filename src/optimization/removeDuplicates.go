/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package optimization

import "main/hashing"

func RemoveDuplicates(arr []string) (result []string) {
	keys := make(map[string]bool);

	for _, value := range arr {
		if _, stored := keys[value]; !stored {
			keys[value] = true;
			result = append(result, value);
		}
	}

	return result;
}

func RemoveEmptyHashes(input []string, algorithm string) (result []string) {
	empty := hashing.Hash("", algorithm);
	space := hashing.Hash(" ", algorithm);

	for _, hash := range input {
		if (hash != empty && hash != space) {
			result = append(result, hash);
		}
	}

	return result;
}