package util

// GetIntegerLength
//  @Description: Get string length of integer
//  @param number: The number need to be measure
//  @return int: The length of number str
func GetIntegerLength(number int) int {
	length := 0
	for number != 0 {
		number /= 10
		length++
	}
	return length
}
