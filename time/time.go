package time

import "time"

func Nanoseconds() int {
	return int(time.Now().UnixNano());
}

func Milliseconds() int {
	return int(time.Now().UnixNano() / 1000000);
}

func Seconds() int {
	return int(time.Now().UnixNano() / 1000000000);
}