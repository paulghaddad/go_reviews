package iteration

// Repeat creates a string repeated 5 times
func Repeat(char string, iterations int) string {
	output := ""
	for i := 0; i < iterations; i++ {
		output += char
	}

	return output
}
