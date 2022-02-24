/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package optimization

import "main/hashing"

/*
* This just goes through and removes duplicate entries from an array so we don't go through the same dictionary-
	entry twice (Or more)
**/
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

/* 
* Empty hashes like spaces and just blank hashes are annoying since we still have to iterate over them so we remove-
	them here
**/
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