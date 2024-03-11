package utils

import (
	"regexp"
	"strings"
)

func ToCamelCase(str string) string {
	words := make([]string, 10)
	re := regexp.MustCompile(`[a-z]+|[A-Z][^A-Z]*`)
	words = re.FindAllString(str, -1)
	// fmt.Println("单词为：", words)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	result := strings.Join(words, "_")
	// fmt.Println("单词最后为",result)
	return result
}
