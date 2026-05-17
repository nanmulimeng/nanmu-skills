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

	// Dynamic thresholds: short queries need lower bars
	thresholdHigh := ThresholdHigh
	minHits := 2
	if len(tokens) <= 3 {
		thresholdHigh = 3.0
		minHits = 1
	}

	type entry struct {
		skill *registry.Skill
		score float64
		hits  int
	}

	var scored []entry
	for i := range m.Skills {
		sk := &m.Skills[i]
		s, h := ScoreWithHits(sk, tokens)
		scored = append(scored, entry{skill: sk, score: s, hits: h})
	}

	// First pass: must requires high score AND sufficient token hits
	for _, e := range scored {
		if e.score >= thresholdHigh && e.hits >= minHits {
			result.Must = append(result.Must, e.skill.Slug)
			mustSet[e.skill.Slug] = true
		}
	}

	// Second pass: remaining skills
	for _, e := range scored {
		if mustSet[e.skill.Slug] {
			continue
		}
		if e.score >= thresholdHigh && e.hits >= minHits {
			result.Must = append(result.Must, e.skill.Slug)
			mustSet[e.skill.Slug] = true
		} else if e.score >= ThresholdLow {
			result.Conditional = append(result.Conditional, e.skill.Slug)
		} else {
			result.Skip = append(result.Skip, e.skill.Slug)
		}
	}

	// Resolve dependency chains
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
