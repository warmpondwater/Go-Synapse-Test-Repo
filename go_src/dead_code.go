package main

// TotallyUnusedGlobalVariable is never read or written
var TotallyUnusedGlobalVariable string = "I am isolated"

// ForgottenFunctionThatDoesNothing is never called by anything
func ForgottenFunctionThatDoesNothing() {
	var localIsolated int = 42
	_ = localIsolated
}

// AbandonedStruct is another isolated struct
type AbandonedStruct struct {
	UselessField int
}
