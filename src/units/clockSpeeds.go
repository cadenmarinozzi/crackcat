/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package units

import (
	"strconv"
	"math"
)

var clockSpeeds = map[string]int{
	"Hz": 1,
	"KHz": 1000,
	"MHz": 1000000,
	"GHz": 1000000000,
};

func ClockSpeedLowerUnit(clockSpeed int, n int) int {
	return clockSpeed / int(math.Pow(1000, float64(n)));
}

func ClockSpeedUpperUnit(clockSpeed int, n int) int {
	return clockSpeed * int(math.Pow(1000, float64(n)));
}

func FormatClockSpeed(clockSpeed int) string {
	lastUnit := "Hz";
	lastSpeed := 1;

	for unit, speed := range clockSpeeds {
		if (clockSpeed > speed) {
			lastUnit = unit;
			lastSpeed = speed;

			continue;
		}
			
		break;
	}
	
	return strconv.Itoa(clockSpeed / lastSpeed) + " " + lastUnit;
}