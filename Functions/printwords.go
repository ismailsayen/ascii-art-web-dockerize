package asciiartweb

func PrintWords(words []string, slice [][]string) string {
	str := ""
	for _, w := range words {
		for i := 1; i <= 8; i++ {
			if len(w) == 0 {
				str += "\r\n"
				break
			}
			for _, e := range w {
				if int(e)-32 >= 0 && int(e)-32 <= len(slice)-1 {
					str += slice[int(e)-32][i]
				} else {
					return "Special charactere is not allowed."
				}
			}
			if i < 8 {
				str += "\r\n"
			}
		}
	}
	if !ContainChars(str) {
		str = str[:len(str)-1]
	}

	return str
}
