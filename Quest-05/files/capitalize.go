package piscine

func Capitalize(s string) string {
	array := []rune(s)
	counter := 0
	for index := range array {
		index = index
		counter++
	}
	if array[0] >= 'a' && array[0] <= 'z' {
		// between 'a' and 'A' there are 32 symbols
		array[0] = array[0] - 32
	}
	for i := 1; i < counter; i++ {
		//decapitalize alphanumeric values
		if array[i] >= 'A' && array[i] <= 'Z' {
			if array[i-1] >= 'A' && array[i-1] <= 'Z' {
				array[i] = array[i] + 32
			}
			if array[i-1] >= 'a' && array[i-1] <= 'z' {
				array[i] = array[i] + 32
			}
			if array[i-1] >= '0' && array[i-1] <= '9' {
				array[i] = array[i] + 32
			}
		}
		if array[i] >= 'a' && array[i] <= 'z' {
			if array[i-1] >= 'A' && array[i-1] <= 'Z' {
				continue
			}
			if array[i-1] >= 'a' && array[i-1] <= 'z' {
				continue
			}
			if array[i-1] >= '0' && array[i-1] <= '9' {
				continue
			} else {
				array[i] = array[i] - 32
			}
		}
	}
	capString := string(array)
	return capString
}
