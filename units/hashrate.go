package units

import "strconv"

func HashRate(hashRate int) string {
	unit := "MH/s";

	if (hashRate <= 1) {
		unit = "KH/s";
		hashRate /= 1000;
	}

	return strconv.Itoa(hashRate) + " " + unit;
}