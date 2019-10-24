package piscine

func Join(strs []string, sep string) string {
	//empty string
	strJoin := ""
	for index, str := range strs {
		//except for the last element always add separator
		if index > 0 {
			strJoin = strJoin + sep + str
		} else {
			strJoin = strJoin + str
		}
	}
	return strJoin
}
