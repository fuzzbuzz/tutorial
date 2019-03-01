package tutorial

// FuzzerEntrypoint is the method Fuzzbuzz will repeatedly
// run with new tests to try and find bugs in BrokenMethod
func FuzzerEntrypoint(data []byte) int {
	// Convert data to the type we need:
	testString := string(data)
	BrokenMethod(testString)
	return 0
}
