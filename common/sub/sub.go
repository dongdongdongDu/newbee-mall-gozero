package sub

// SubStrLen 超过指定长度后显示...
func SubStrLen(str string, length int) string {
	nameRune := []rune(str)
	if len(nameRune) > length {
		return string(nameRune[:length-1]) + "..."
	}

	return str
}
