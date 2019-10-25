package piscine

func Index(s string, toFind string) int {
	var comp string
	for i := 0; i < StrLen(s)-StrLen(toFind); i++ {
		for j := 0; j < StrLen(toFind); j++ {
			comp = comp + string(s[i+j])
		}
		if comp == toFind {
			return i
		}
		comp = ""
	}
	return -1
}
