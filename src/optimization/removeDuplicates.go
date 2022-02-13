/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package optimization

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