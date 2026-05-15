package registry

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func LoadFromDir(dir string) (*Manifest, error) {
	var skills []Skill

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "SKILL.md" {
			sk, err := parseSkillFile(path)
			if err == nil && sk.Slug != "" {
				skills = append(skills, *sk)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return NewManifest(skills), nil
}

func parseSkillFile(path string) (*Skill, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sk := &Skill{
		Slug: filepath.Base(filepath.Dir(path)),
	}

	scanner := bufio.NewScanner(f)
	inFrontmatter := false
	fmLines := []string{}
	bodyLines := []string{}
	afterFM := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				inFrontmatter = false
				afterFM = true
				continue
			}
		}
		if inFrontmatter {
			fmLines = append(fmLines, line)
		} else if afterFM {
			bodyLines = append(bodyLines, line)
		}
	}

	parseFrontmatter(sk, fmLines)
	sk.Body = strings.TrimSpace(strings.Join(bodyLines, "\n"))

	if sk.Name == "" {
		sk.Name = sk.Slug
	}

	if len(sk.Triggers) == 0 {
		sk.Triggers = extractKeywords(sk.Description, sk.Slug)
	}

	return sk, nil
}

func parseFrontmatter(sk *Skill, lines []string) {
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "name":
			sk.Name = unquote(val)
		case "description":
			sk.Description = unquote(val)
		case "category":
			sk.Category = unquote(val)
		case "triggers":
			sk.Triggers = parseList(val)
		case "requires":
			sk.Requires = parseList(val)
		case "related":
			sk.Related = parseList(val)
		case "version":
			sk.Version = unquote(val)
		}
	}
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	if len(s) >= 2 && s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	}
	return s
}

func parseList(s string) []string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '[' && s[len(s)-1] == ']' {
		s = s[1 : len(s)-1]
	}
	parts := strings.Split(s, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		p = unquote(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func extractKeywords(description, slug string) []string {
	seen := make(map[string]bool)
	var kw []string

	add := func(w string) {
		w = strings.ToLower(strings.TrimSpace(w))
		if len(w) >= 2 && !seen[w] && !isStopWord(w) {
			seen[w] = true
			kw = append(kw, w)
		}
	}

	// Split by common Chinese/English separators
	seps := []string{" - ", " -", "- ", "、", "，", "。", "；", "：", "/", "|", " ", ","}
	parts := []string{description}
	for _, sep := range seps {
		var next []string
		for _, p := range parts {
			next = append(next, strings.Split(p, sep)...)
		}
		parts = next
	}

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			add(p)
		}
	}

	// Add slug as trigger
	add(slug)

	// Extract bigrams from Chinese text
	chinese := filterChinese(description)
	for i := 0; i+2 <= len([]rune(chinese)); i++ {
		bg := string([]rune(chinese)[i : i+2])
		if len(bg) == 6 { // 2 Chinese chars = 6 bytes in UTF-8
			add(bg)
		}
	}

	return kw
}

func filterChinese(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 0x4e00 && r <= 0x9fff {
			result = append(result, r)
		}
	}
	return string(result)
}

func isStopWord(w string) bool {
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
