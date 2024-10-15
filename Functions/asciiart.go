package asciiartweb

import (
	"bufio"
	"os"
)

func AsciiArt(sentence, banner string) string {
	file, err := os.Open("./Files/" + banner + ".txt")
	if err != nil {
		return "Banner not found"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	symbole := []string{}
	symboles := [][]string{}

	for scanner.Scan() {
		symbole = append(symbole, scanner.Text())
		count++

		if count == 9 {
			symboles = append(symboles, symbole)
			symbole = []string{}
			count = 0
		}
	}
	if len(symboles) < 95 {
		return "All caracters"
	}
	return PrintWords(Split(sentence), symboles)
}
