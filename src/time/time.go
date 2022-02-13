/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package time

import "time"

func Nanoseconds() int {
	return int(time.Now().UnixNano());
}

func Milliseconds() int {
	return int(Nanoseconds() / 1000000);
}

func Seconds() int {
	return int(Milliseconds() / 1000);
}