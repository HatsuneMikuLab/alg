package zalg

func matchPattern(zArr []int, source string, index int, right int, left int) int {
	for right < len(source) && source[right-left] == source[right] {
		right++
	}
	zArr[index] = right - left
	return right - 1
}

func CalcZValues(source string) []int {
	//patterLength := 0
	left := 0
	right := 0
	index := 1
	result := make([]int, len(source))
	for index < len(source) {
		if right < index {
			left, right = index, index
			right = matchPattern(result, source, index, right, left)
		} else {
			k := index - left
			if result[k] < right-index+1 {
				result[index] = result[k]
			} else {
				left = index
				right = matchPattern(result, source, index, right, left)
			}
		}
		index++
	}
	return result
}
