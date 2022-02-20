/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package rules

import (
	"strings"
	"strconv"
)

func GenerateRules(dictionary []string, rulesList []string) (newDictionary []string) {
	var memory string;

	for _, entry := range dictionary {
		if (entry == "" || len(entry) == 0) { continue }

		for _, rules := range rulesList {
			if (rules == "" || rules == "\n") { continue }
			if (len(rules) >= 2 && string(rules[:2]) == "##") { continue }
			
			newEntry := entry;
			position := 0;

			for (position < len(rules)) {
				rule := string(rules[position]);
				newEntryLength := len(newEntry);

				if (newEntry == "" || newEntryLength == 0) { continue }

				switch (rule) {
					case ":", "":
						break;

					case ("l"): // Lowercase
						newEntry = strings.ToLower(newEntry);

						break;
	
					case ("u"): // Uppercase
						newEntry = strings.ToUpper(newEntry);

						break;
	
					case ("c"): // Upper first char, lower the rest
						firstCharacter := string(newEntry[0]);
						newEntry = strings.ToUpper(firstCharacter) + strings.ToLower(newEntry[1:]);

						break;
	
					case ("C"): // Lower first char, upper the rest
						firstCharacter := string(newEntry[0]);
						newEntry = strings.ToLower(firstCharacter) + strings.ToUpper(newEntry[1:]);

						break;
	
					case ("t"): // Toggle case
						newEntry = ToggleCase(newEntry);

						break;
	
					case ("T"): // Toggle case at index
						position++;
						index, _ := strconv.Atoi(string(rules[position]));
						ToggleCaseAtX(newEntry, index);

						break;

					case ("r"): // Reverse
						newEntry = ReverseString(newEntry);

						break;

					case ("d"): // Duplicate
						newEntry += newEntry;

						break;

					case ("p"): // Duplicate n times
						position++;
						n, _ := strconv.Atoi(string(rules[position]));
						tempEntry := newEntry;

						for i := 0; i < n; i++ {
							newEntry += tempEntry;
						}

						break;

					case ("f"): // Mirror
						newEntry += ReverseString(newEntry);

						break;

					case ("{"): // Append first character to end
						newEntry = newEntry[1:] + string(newEntry[0]);

						break;

					case ("}"): // Append last character to start
						newEntry = string(newEntry[newEntryLength - 1]) + newEntry[:newEntryLength - 1];

						break;
				
					case ("$"): // Append x to end
						position++;
						newEntry += string(rules[position]);

						break;

					case ("^"): // Append x to start
						position++;
						newEntry = string(rules[position]) + newEntry;

						break;

					case ("["): // Remove first character
						newEntry = newEntry[1:newEntryLength];

						break;

					case ("]"): // Remove last character
						newEntry = newEntry[:newEntryLength - 1];

						break;

					case ("D"): // Remove character at n
						position++;
						n, _ := strconv.Atoi(string(rules[position]));

						newEntry = ReplaceCharacterAtX(newEntry, "", n);

						break;

					case ("x"): // Get the characters between N and M
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						position += 2;

						if (newEntryLength >= N && newEntryLength >= M) {
							newEntry = newEntry[N:M];
						}

						break;

					case ("O"): // Remove all characters between N and M
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						position += 2;

						if (newEntryLength >= N && newEntryLength >= N + M) {
							newEntry = newEntry[:N] + newEntry[N + M:];
						}

						break;

					case ("i"): // Insert x at N
						N, _ := strconv.Atoi(string(rules[position + 1]));
						X := string(rules[position + 2]);
						position += 2;

						if (newEntryLength >= N) {
							newEntry = newEntry[:N] + X + newEntry[N:];
						}

						break;

					case ("o"): // Replace character at n with x
						N, _ := strconv.Atoi(string(rules[position + 1]));
						X := string(rules[position + 2]);
						position += 2;

						newEntry = ReplaceCharacterAtX(newEntry, X, N);

						break;

					case ("'"): // Remove all characters after n
						position++;
						N, _ := strconv.Atoi(string(rules[position]));
						
						if (newEntryLength >= N) {
							newEntry = newEntry[:N];
						}

						break;

					case ("s"): // Replace x with y
						X := rules[position + 1];
						Y := rules[position + 2];
						position += 2;
						tempEntry := []byte(newEntry);
						
						for i, character := range tempEntry {
							if (character == X) {
								tempEntry[i] = Y;
							}
						}

						newEntry = string(tempEntry);

						break;

					case ("@"): // Remove x
						position++;
						X := rules[position];
						tempEntry := []byte(newEntry);
						
						for i, character := range tempEntry {
							if (character == X) {
								tempEntry[i] = byte(0);
							}
						}

						newEntry = string(tempEntry);

						break;

					case ("z"): // Duplicate first character n times
						position++;
						N, _ := strconv.Atoi(string(rules[position]));
						tempEntry := "";

						for i := 0; i < N; i++ {
							tempEntry += string(newEntry[0]);
						}

						newEntry += tempEntry + newEntry[1:];

						break;

					case ("Z"):// Duplicate last character n times
						position++;
						N, _ := strconv.Atoi(string(rules[position]));

						for i := 0; i < N; i++ {
							newEntry += string(newEntry[newEntryLength - 1]);
						}
						
						break;

					case ("q"): // Duplicate all characters
						tempEntry := "";

						for _, character := range newEntry {
							characterString := string(character);
							tempEntry += characterString + characterString;
						}

						newEntry = tempEntry;
						
						break;

					case ("X"): // Insert to memory
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						I, _ := strconv.Atoi(string(rules[position + 3]));
						position += 3;
						mem := memory[N:N + M];
						newEntry = newEntry[:I] + mem + newEntry[I:newEntryLength];
						
						break;

					case ("4"): // Append memory to end
						newEntry += memory;
						
						break;

					case ("6"): // Append memory to start
						newEntry = memory + newEntry;
						
						break;

					case ("M"): // Set memory
						memory = newEntry;
						
						break;

					default:					
						break;
				}

				position++;
			}

			newDictionary = append(newDictionary, newEntry);
		}
	}

	return newDictionary;
}