package iteration1

func Repeat(char string, times int) string {
	output := ""
	for i := 0; i < times; i++ {
		output += char
	}

	return output
}
