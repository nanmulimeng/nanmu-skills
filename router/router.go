package router

import "nanmu-skill-mcp/registry"

type RouteResult struct {
	Must        []string `json:"must"`
	Conditional []string `json:"conditional"`
	Skip        []string `json:"skip"`
}

func Route(m *registry.Manifest, request string) *RouteResult {
	tokens := Tokenize(request)
	mustSet := make(map[string]bool)
	result := &RouteResult{}

	var scored []ScoredSkill
	for i := range m.Skills {
		sk := &m.Skills[i]
		s := Score(sk, tokens, mustSet)
		scored = append(scored, ScoredSkill{Skill: sk, Score: s})
	}

	// First pass: determine must based on high scores
	for _, ss := range scored {
		if ss.Score >= ThresholdHigh {
			result.Must = append(result.Must, ss.Skill.Slug)
			mustSet[ss.Skill.Slug] = true
		}
	}

	// Second pass: re-score with updated mustSet for dependency boosts
	result.Conditional = nil
	result.Skip = nil
	for _, ss := range scored {
		if mustSet[ss.Skill.Slug] {
			continue
		}
		s := Score(ss.Skill, tokens, mustSet)
		if s >= ThresholdHigh {
			result.Must = append(result.Must, ss.Skill.Slug)
			mustSet[ss.Skill.Slug] = true
		} else if s >= ThresholdLow {
			result.Conditional = append(result.Conditional, ss.Skill.Slug)
		} else {
			result.Skip = append(result.Skip, ss.Skill.Slug)
		}
	}

	// Resolve dependencies
	resolveDeps(m, result, mustSet)

	return result
}

func resolveDeps(m *registry.Manifest, result *RouteResult, mustSet map[string]bool) {
	allSelected := make(map[string]bool)
	for _, s := range result.Must {
		allSelected[s] = true
	}
	for _, s := range result.Conditional {
		allSelected[s] = true
	}

	queue := append([]string{}, result.Must...)
	queue = append(queue, result.Conditional...)

	for len(queue) > 0 {
		slug := queue[0]
		queue = queue[1:]

		sk := m.Get(slug)
		if sk == nil {
			continue
		}
		for _, req := range sk.Requires {
			if !allSelected[req] {
				result.Must = append(result.Must, req)
				allSelected[req] = true
				mustSet[req] = true
				queue = append(queue, req)
			}
		}
	}
}
