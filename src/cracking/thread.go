/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package cracking

// type callback func(*Thread)

type Thread struct {
	Index int;
	EntryIndex int;
	EndIndex int;
	
	Running bool;
	// Callback callback;
}