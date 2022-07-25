package utils

func NullValueFromType(typeKey *string) string {
	if *typeKey == "INT" {
		return "0"
	}
	if *typeKey == "BOOL" {
		return "FALSE"
	}
	return "''"
}

func RemoveLastNChars(str *string, n int) {
	*str = string([]rune(*str)[:len(*str)-n])
}
func RemoveFirstNChars(str string, n int) string {
	return str[n:]
}
