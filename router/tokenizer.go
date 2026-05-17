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

	// Extract alphanumeric sequences from mixed text ("部署K8s" → "k8s")
	var buf strings.Builder
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			buf.WriteRune(r)
		} else {
			if buf.Len() >= 2 {
				add(buf.String())
			}
			buf.Reset()
		}
	}
	if buf.Len() >= 2 {
		add(buf.String())
	}

	// Extract English words from whitespace/punctuation splits
	words := strings.Fields(text)
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w == "" {
			continue
		}
		clean := strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) || unicode.IsSymbol(r) ||
				r == '，' || r == '。' || r == '！' || r == '？' ||
				r == '；' || r == '：' || r == '、' {
				return ' '
			}
			return r
		}, w)
		for _, part := range strings.Fields(clean) {
			if isAscii(part) && len(part) >= 2 {
				add(part)
			}
		}
	}

	// Chinese bigrams only (no single chars — too noisy for scoring)
	chinese := extractChinese(text)
	runes := []rune(chinese)
	for i := 0; i+1 < len(runes); i++ {
		bg := string(runes[i : i+2])
		add(bg)
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
		"me": true, "my": true, "no": true, "do": true, "if": true,
	}
	return stop[w]
}
