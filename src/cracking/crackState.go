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
	
	Algorithm string;
	Threads int;
	MaxTime int;

	CrackingMethod string;
	CrackingMode string;

	Found []string;

	Iterations int;
	StartTime int;
	EndTime int;
	NPasswords int;

	RemoveFound bool;
	LogFound bool;
	SameLineLogs bool;

	FormattedStartTime string;
	FormattedEndTime string;
}