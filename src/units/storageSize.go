/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package units

import (
	"strconv"
	"math"
)

var storageSizes = map[string]int{
	// "bits": 1,
	// "bytes": 8,
	"KB": 1000,
	"MB": 1000000,
	"GB": 1000000000,
	"TB": 1000000000000,
};

func StorageSizeLowerUnit(storageSize int, n int) int {
	return storageSize / int(math.Pow(1000, float64(n)));
}

func StorageSizeUpperUnit(storageSize int, n int) int {
	return storageSize * int(math.Pow(1000, float64(n)));
}


func FormatStorageSize(storageSize int) string {
	lastUnit := "KB";
	lastSize := 1000;

	for unit, size := range storageSizes {
		if (storageSize > size) {
			lastUnit = unit;
			lastSize = size;

			continue;
		}
			
		break;
	}
	
	return strconv.Itoa(storageSize / lastSize) + " " + lastUnit;
}