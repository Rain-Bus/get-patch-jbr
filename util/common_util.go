package util

func GetIntegerLength(number int) int {
	length := 0
	for number != 0 {
		number /= 10
		length++
	}
	return length
}
