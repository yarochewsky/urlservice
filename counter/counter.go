package counter

import (
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func MapFrequencies(document string) map[string]int {
	m := make(map[string]int)

	tokenizer := html.NewTokenizer(strings.NewReader(document))

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			break
		}
		tag := tokenizer.Text()

		parts := strings.Split(strings.ToLower(string(tag)), " ")
		for _, p := range parts {
			// remove word separators
			pattern := regexp.MustCompile(`[.:,"\n\t]+`)
			pp := pattern.ReplaceAllString(p, "")

			if pp == "" {
				// remove empty words as a result of separation or otherwise
				continue
			}

			if _, ok := m[pp]; !ok {
				m[pp] = 0
			}
			m[pp] += 1
		}
	}

	return m
}
