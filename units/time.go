package units

import "strconv"

func LongHandTime(time int) string {
	unit := "milliseconds";

	if (time >= 1000) {
		unit = "seconds";

		if (time <= 2000) {
			unit = "second";
		}

		time /= 1000;
	}

	return strconv.Itoa(time) + " " + unit;
}

func ShortHandTime(time int) string {
	unit := "ms";

	if (time >= 1000) {
		unit = "s";
		time /= 1000;
	}

	return strconv.Itoa(time) + unit;
}