package piscine

func ConcatParams(args []string) string {
	//empty string
	str := ""
	//lenght of an args array
	var len int
	for index := range args {
		len = index + 1
	}
	for index, element := range args {
		str += element
		if index != len-1 {
			str += "\n"
		}
	}
	return str
}
