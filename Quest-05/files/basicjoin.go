package piscine

func BasicJoin(strs []string) string {
	//empty string
	strJoin := ""
	for _, element := range strs {
		strJoin = strJoin + element
	}
	return strJoin
}
