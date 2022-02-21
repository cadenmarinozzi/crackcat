/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package terminal

import (
	"main/units"
	"main/cracking"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"strings"
)

// This is just all the stuff that is printed to the console (Besides realtime stuff)

const WATERMARK = "                                        /$$                             /$$    \n                                       | $$                            | $$    \n  /$$$$$$$  /$$$$$$  /$$$$$$   /$$$$$$$| $$   /$$  /$$$$$$$  /$$$$$$  /$$$$$$  \n /$$_____/ /$$__  $$|____  $$ /$$_____/| $$  /$$/ /$$_____/ |____  $$|_  $$_/  \n| $$      | $$  \\__/ /$$$$$$$| $$      | $$$$$$/ | $$        /$$$$$$$  | $$    \n| $$      | $$      /$$__  $$| $$      | $$_  $$ | $$       /$$__  $$  | $$ /$$\n|  $$$$$$$| $$     |  $$$$$$$|  $$$$$$$| $$ \\  $$|  $$$$$$$|  $$$$$$$  |  $$$$/\n \\_______/|__/      \\_______/ \\_______/|__/  \\__/ \\_______/ \\_______/   \\___/\n";
const KITTY = "       _                        \n       \\`*-.                    \n        )  _`-.                 \n       .  : `. .                \n       : _   '  \\               \n       ; *` _.   `*-._          \n       `-.-'          `-.       \n         ;       `       `.     \n         :.       .        \\    \n         . \\  .   :   .-'   .   \n         '  `+.;  ;  '      :   \n         :  '  |    ;       ;-. \n         ; '   : :`-:     _.`* ;\n      .*' /  .*' ; .*`- +'  `*' \n     `*-*   `*-*  `*-*'";

func Watermark(kitty bool) {
	fmt.Println(WATERMARK);

	if (kitty) {
		fmt.Println(KITTY);
	}

	fmt.Println("\t\tby: nekumelon\n");
}

func Header(version string) {
	fmt.Printf("crackcat (v%s) starting...\n", version);
}

func Devices() {
	CPU, _ := cpu.Info();
	cpuInfo := CPU[0];
	cores := int(cpuInfo.Cores);
	clockSpeed := units.ClockSpeedUpperUnit(int(cpuInfo.Mhz), 2);
	cacheSize := units.StorageSizeUpperUnit(int(cpuInfo.CacheSize), 2);
	cpuName := strings.Split(cpuInfo.ModelName, "           ")[0];

	fmt.Printf("* CPU: %s, %s, %d cores, cache size: %s\n\n", cpuName, units.FormatClockSpeed(clockSpeed), cores, units.FormatStorageSize(cacheSize));
}

func Optimizers(optimizers map[string]bool) {
	fmt.Println("Optimizers:");
	
	for optimizer, enabled := range optimizers {
		if (enabled) {
			fmt.Printf("* %s\n", optimizer);
		}
	}

	fmt.Println();
}

func Cracked(state cracking.CrackState) {
	foundPercent := int((float64(len(state.Found)) / float64(state.NPasswords)) * 100);
	fmt.Printf("Start time......: %d (%s)\n", state.StartTime, state.FormattedStartTime);
	fmt.Printf("End time........: %d (%s)\n", state.EndTime, state.FormattedEndTime);
	fmt.Printf("Time taken......: %d seconds\n", state.EndTime - state.StartTime);
	fmt.Printf("Speed...........: %d passwords per second\n", len(state.Found) / state.MaxTime);
	fmt.Printf("Entries speed...: %d entries per second\n", state.Iterations / state.MaxTime);
	fmt.Printf("Iterations......: %d\n", state.Iterations);
	fmt.Printf("Found...........: %d passwords (%d%% of %d passwords)\n", len(state.Found), foundPercent, state.NPasswords);
	fmt.Printf("Algorithm.......: %s", state.Algorithm);
}