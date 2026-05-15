package router

import "nanmu-skill-mcp/registry"

const (
	ThresholdHigh = 2.0
	ThresholdLow  = 0.5
)

type ScoredSkill struct {
	Skill *registry.Skill
	Score float64
}

func Score(sk *registry.Skill, tokens []string, mustSet map[string]bool) float64 {
	score := 0.0

	for _, token := range tokens {
		for _, trigger := range sk.Triggers {
			if token == trigger || containsWord(trigger, token) {
				score += 2.0
				break
			}
		}
		if containsWord(sk.Description, token) {
			score += 1.0
		}
		if containsWord(sk.Name, token) {
			score += 1.5
		}
	}

	// Dependency boost
	for _, req := range sk.Requires {
		if mustSet[req] {
			score *= 1.5
			break
		}
	}
	for _, rel := range sk.Related {
		if mustSet[rel] {
			score *= 1.2
			break
		}
	}

	return score
}

func containsWord(s, word string) bool {
	return len(s) > 0 && len(word) > 0 && len(s) >= len(word) &&
		findWord(s, word)
}

func findWord(s, word string) bool {
	for i := 0; i <= len(s)-len(word); i++ {
		if s[i:i+len(word)] == word {
			return true
		}
	}
	return false
}
