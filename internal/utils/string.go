package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToPascal(i string) string {
	caser := cases.Title(language.English)

	words := strings.Split(i, "-")

	text := ""
	for _, word := range words {
		text += caser.String(word)
	}

	return text
}
