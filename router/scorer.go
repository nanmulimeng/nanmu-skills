package router

import "nanmu-skill-mcp/registry"

const (
	ThresholdHigh = 6.0
	ThresholdLow  = 2.0
)

type ScoredSkill struct {
	Skill *registry.Skill
	Score float64
}

// ScoreWithHits returns score and distinct token hit count.
func ScoreWithHits(sk *registry.Skill, tokens []string) (float64, int) {
	score := 0.0
	hits := 0

	for _, token := range tokens {
		matched := false

		// Exact trigger match (strongest signal)
		for _, trigger := range sk.Triggers {
			if token == trigger {
				score += 3.0
				matched = true
				hits++
				break
			}
		}
		if matched {
			continue
		}

		// Trigger substring match (weak — bigrams overlap)
		for _, trigger := range sk.Triggers {
			if containsWord(trigger, token) {
				score += 0.5
				matched = true
				hits++
				break
			}
		}
		if matched {
			continue
		}

		// Description match (weakest — descriptions are long)
		if containsWord(sk.Description, token) {
			score += 0.3
			hits++
			continue
		}

		// Name match
		if containsWord(sk.Name, token) {
			score += 0.5
			hits++
		}
	}

	return score, hits
}

// Score returns only the score (for compatibility).
func Score(sk *registry.Skill, tokens []string, _ map[string]bool) float64 {
	s, _ := ScoreWithHits(sk, tokens)
	return s
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
