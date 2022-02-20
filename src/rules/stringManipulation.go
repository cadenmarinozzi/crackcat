/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package rules

import "unicode"

func ReverseString(input string) (reversed string) {
	for _, character := range input {
		reversed = string(character) + reversed;
	}

	return reversed;
}

func ReplaceCharacterAtX(input string, replacement string, x int) string {
	inputLength := len(input) - 1;

	if (inputLength > x) {
		return input[:x] + replacement + input[x + 1:];
	} else if (inputLength == x) {
		return input[:x] + replacement;
	}

	return input;
}

func ToggleCase(input string) string {
	tempInput := []rune(input)

	for i, char := range tempInput {
		if unicode.IsUpper(char) {
			tempInput[i] = unicode.ToLower(char)

			continue
		}

		tempInput[i] = unicode.ToUpper(char)
	}

	return string(tempInput)
}

func ToggleCaseAtX(input string, x int) string {
	if (len(input) - 1 >= x) {
		tempInput := []rune(input);
		character := tempInput[x];

		if (unicode.IsUpper(character)) {
			tempInput[x] = unicode.ToLower(character);
		} else {
			tempInput[x] = unicode.ToUpper(character);
		}

		return string(tempInput);
	}

	return input;
}