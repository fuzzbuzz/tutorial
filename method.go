package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of data even when it only has a length of 3.
func BrokenMethod(data string) bool {
	return len(data) >= 3 &&
		data[0] == 'F' &&
		data[1] == 'U' &&
		data[2] == 'Z' &&
		data[3] == 'Z'
}
