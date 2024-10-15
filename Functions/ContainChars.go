package asciiartweb

func ContainChars(s string) bool {
	// verifie si le string contient des char ou juste les \n
	for _, r := range s {
		if int(r) >= 32 && int(r) <= 126 {
			return true
		}
	}
	return false
}
