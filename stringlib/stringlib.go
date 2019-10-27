package stringlib

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func substr(s string, index int, size int) string {
	if len(s) < index+size {
		return s[index:]
	}
	return s[index : index+size]
}
