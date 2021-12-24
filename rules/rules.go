package rules

import (
	"strings"
	"bytes"
	"unicode"
	"strconv"
)

func reverse(str string) (result string) {
    for _, character := range str {
        result = string(character) + result;
    }

	return result;
}

var memory string;

func GenDictionary(dictionary []string, rulesList []string) (outputDictionary []string) {
	for _, password := range dictionary {
		for _, rules := range rulesList {
			if (rules == "") { continue }
			if (len(rules) > 1 && string(rules[:2]) == "##") { continue }

			position := 0;
			newPassword := password;

			for (position < len(rules)) {
				rule := string(rules[position]);

				switch (rule) {
					case (":"):
						newPassword = newPassword;

					case ("l"):
						newPassword = strings.ToLower(newPassword);

					case ("u"):
						newPassword = strings.ToUpper(newPassword);

					case ("c"):
						newPassword = strings.ToUpper(string(newPassword[0])) + strings.ToLower(string(newPassword[1:len(newPassword)]));

					case ("C"):
						newPassword = strings.ToLower(string(newPassword[0])) + strings.ToUpper(string(newPassword[1:len(newPassword)]));
					
					case ("t"):
						temp := []byte(newPassword);

						for i := 0; i < len(temp); i++ {
							tempRune := rune(temp[i]);

							if (unicode.IsUpper(tempRune)) {
								temp[i] = bytes.ToLower(temp)[i];
							} else if (unicode.IsLower(tempRune)) {
								temp[i] = bytes.ToUpper(temp)[i];
							}
						}

						newPassword = string(temp);

					case ("T"):
						position++;
						index, _ := strconv.Atoi(string(rules[position]));

						if (len(newPassword) >= index) {
							temp := []byte(newPassword);
							tempRune := rune(temp[index]);

							if (unicode.IsUpper(tempRune)) {
								temp[index] = bytes.ToLower(temp)[index];
							} else if (unicode.IsLower(tempRune)) {
								temp[index] = bytes.ToUpper(temp)[index];
							}

							newPassword = string(temp);
						}

					case ("r"):
						newPassword = reverse(newPassword);

					case ("d"):
						newPassword = newPassword + newPassword;

					case ("p"):
						position++;
						index, _ := strconv.Atoi(string(rules[position]));

						for i := 0; i < index; i++ {
							newPassword += newPassword;
						}

					case ("f"):
						newPassword = newPassword + reverse(newPassword);

					case ("{"):
						newPassword = newPassword[1:len(newPassword)] + string(newPassword[0]);

					case ("}"):
						newPassword = string(newPassword[len(newPassword) - 1]) + newPassword[:len(newPassword)];
				
					case ("$"):
						position++;
						newPassword = newPassword + string(rules[position]);

					case ("^"):
						position++;
						newPassword = string(rules[position]) + newPassword;

					case ("["):
						newPassword = newPassword[1:len(newPassword)];

					case ("]"):
						newPassword = newPassword[:len(newPassword) - 1];

					case ("D"):
						position++;
						index, _ := strconv.Atoi(string(rules[position]));
						
						for i, character := range newPassword {
							if (i != index) {
								newPassword += string(character);
							}
						}

					case ("x"):
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						position += 2;
						newPassword = newPassword[N:M];

					case ("O"):
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						position += 2;
						
						for i, character := range newPassword {
							if (i < N || i >= N + M) {
								newPassword += string(character);
							}
						}

					case ("i"):
						N, _ := strconv.Atoi(string(rules[position + 1]));
						X := string(rules[position + 2]);
						position += 2;
						
						for i, character := range newPassword {
							if (i == N) {
								newPassword += X;
							}

							newPassword += string(character);
						}

					case ("o"):
						N, _ := strconv.Atoi(string(rules[position + 1]));
						X := string(rules[position + 2]);
						position += 2;
						
						for i, character := range newPassword {
							if (i == N) {
								newPassword += X;

								continue;
							}

							newPassword += string(character);
						}

					case ("\""):
						position++;
						N, _ := strconv.Atoi(string(rules[position]));
						
						for i, character := range newPassword {
							if (i != N) {
								newPassword += string(character);
							}
						}

					case ("s"):
						X := string(rules[position + 1]);
						Y := string(rules[position + 2]);
						position += 2;
						
						for _, character := range newPassword {
							if (string(character) == X) {
								newPassword += Y;

								continue;
							}

							newPassword += string(character);
						}

					case ("@"):
						position++;
						X := string(rules[position]);
						
						for _, character := range newPassword {
							if (string(character) == X) {
								continue;
							}

							newPassword += string(character);
						}

					case ("z"):
						position++;
						N, _ := strconv.Atoi(string(rules[position]));

						for i := 0; i < N; i++ {
							newPassword += string(newPassword[0]);
						}

						newPassword += newPassword;

					case ("Z"):
						position++;
						N, _ := strconv.Atoi(string(rules[position]));
						newPassword += newPassword;

						for i := 0; i < N; i++ {
							newPassword += string(newPassword[len(newPassword) - 1]);
						}

					case ("q"):
						for _, character := range newPassword {
							newPassword += string(character) + string(character);
						}

					case ("X"):
						N, _ := strconv.Atoi(string(rules[position + 1]));
						M, _ := strconv.Atoi(string(rules[position + 2]));
						I, _ := strconv.Atoi(string(rules[position + 3]));
						position += 3;
						mem := memory[N:N + M];
						newPassword = newPassword[:I] + mem + newPassword[I:len(newPassword)];

					case ("4"):
						newPassword = newPassword + memory;

					case ("6"):
						newPassword = memory + newPassword;

					case ("M"):
						memory = newPassword;
				}

				position++;
			}

			outputDictionary = append(outputDictionary, newPassword);
		}
	}

	return outputDictionary;
}