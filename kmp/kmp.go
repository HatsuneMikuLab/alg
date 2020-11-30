package kmp

func BuildLPSArray(pattern string) []int {
	patterLength := 0
	index := 1
	result := make([]int, len(pattern))
	for index < len(pattern) {
		if pattern[index] == pattern[patterLength] {
			patterLength++
			result[index] = patterLength
		} else {
			patterLength = result[patterLength]
		}
		index++
	}
	return result
}

/*func ExecAlg(source string, pattern string) []int {
	offset := 0
	rightPointer := 1
	patternLength := 0
	result := make([]int, len(input))
	for rightPointer < len(input) {
		if input[offset] == input[rightPointer+offset] {
			offset++
			patternLength++
		} else {
			result = append(result, patternLength)
			patternLength = 0
			leftPointer = 0
			rightPointer++
		}
	}
	return result
}
*/
