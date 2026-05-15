package router

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	var tokens []string
	seen := make(map[string]bool)

	add := func(w string) {
		w = strings.ToLower(strings.TrimSpace(w))
		if len(w) >= 2 && !seen[w] && !isStop(w) {
			seen[w] = true
			tokens = append(tokens, w)
		}
	}

	// First, extract pure English/ASCII words split by whitespace
	words := strings.Fields(text)
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w == "" {
			continue
		}
		// Extract English words even from mixed text like "REST API，需要"
		clean := strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) || unicode.IsSymbol(r) ||
				r == '，' || r == '。' || r == '！' || r == '？' ||
				r == '；' || r == '：' || r == '、' {
				return ' '
			}
			return r
		}, w)
		for _, part := range strings.Fields(clean) {
			if isAscii(part) {
				add(part)
			}
		}
	}

	// Extract Chinese bigrams
	chinese := extractChinese(text)
	runes := []rune(chinese)
	for i := 0; i+1 < len(runes); i++ {
		bg := string(runes[i : i+2])
		add(bg)
	}
	for _, r := range runes {
		ug := string(r)
		if !isStop(ug) && !seen[ug] {
			seen[ug] = true
			tokens = append(tokens, ug)
		}
	}

	return tokens
}

func extractChinese(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 0x4e00 && r <= 0x9fff {
			result = append(result, r)
		}
	}
	return string(result)
}

func isAscii(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func isStop(w string) bool {
	stop := map[string]bool{
		"的": true, "是": true, "在": true, "了": true, "和": true,
		"与": true, "或": true, "对": true, "等": true, "及": true,
		"为": true, "被": true, "把": true, "从": true, "让": true,
		"a": true, "an": true, "the": true, "is": true, "are": true,
		"to": true, "of": true, "in": true, "for": true, "and": true,
		"or": true, "it": true, "on": true, "with": true, "as": true,
		"this": true, "be": true, "by": true, "at": true, "你": true,
	}
	return stop[w]
}
