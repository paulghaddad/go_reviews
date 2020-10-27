package iteration2

func Repeat(letter string, times int) string {
	output := ""

	for i := 0; i < times; i++ {
		output += letter
	}

	return string(output)
}
