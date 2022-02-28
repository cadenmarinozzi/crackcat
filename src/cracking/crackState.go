/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package cracking

// This is my favorite file by far

type CrackState struct {
	SessionName string;
	
	Passwords []string;
	Dictionary []string;
	Found []string;
	
	Algorithm string;
	Threads int;
	MaxTime int;
	BenchmarkTime int;
	EstimatedTime int;
	DeltaTime int;

	CrackingMethod string;
	CrackingMode string;

	Iterations int;
	NPasswords int;

	RemoveFound bool;
	LogFound bool;
	SameLineLogs bool;

	FormattedStartTime string;
	FormattedEndTime string;
	StartTime int;
	EndTime int;
}