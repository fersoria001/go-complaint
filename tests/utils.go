package tests

func RepeatString(str string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += str
	}
	return result
}
