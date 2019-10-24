package piscine

func Index(s string, toFind string) int {
	input := []rune(s)
	search := []rune(toFind)
	input_i := 0
	search_i := 0

	for index := range search {
		//to implement range
		index = index
		search_i++
	}
	if search_i == 0 {
		return 0
	}
	for index := range input {
		//to implement range
		index = index
		input_i++
	}
	for index, char := range input {
		if char == search[0] && input_i >= search_i+index-1 {
			var i int = 1
			for j := 1; j < search_i; j++ {
				if search[j] == input[index+j] {
					i++
				}
			}
			if i == search_i {
				return index
			}
		}
	}
	return -1
}
