package basic

// FindFirstRepeatedChar 查找第一个重复的字符
func FindFirstRepeatedChar(chars string) rune {
	m := make(map[rune]int)
	for _, c := range chars {
		m[c] = m[c] + 1
		if m[c] == 2 {
			return c
		}
	}
	return -1
}
