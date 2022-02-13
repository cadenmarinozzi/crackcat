/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package cracking

import (
	"fmt"
	"os"
)

func Crack(state CrackState) CrackState {
	switch (state.CrackingMethod) {
		case ("dictionary"):
			state = DictionaryAttack(state);

			break;

		default:
			fmt.Println("Unsupported cracking method");
			os.Exit(1);

			break;
	}

	return state;
}