package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		//return an array with nil value
		var array []int = nil
		return array
	} else {
		array := make([]int, max-min)
		index := 0
		for i := min; i < max; i++ {
			array[index] = i
			index++
		}
		return array
	}
}
