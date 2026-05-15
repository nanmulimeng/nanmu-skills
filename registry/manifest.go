package registry

import "time"

type Skill struct {
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Triggers    []string `json:"triggers"`
	Requires    []string `json:"requires"`
	Related     []string `json:"related"`
	Body        string   `json:"body"`
	Version     string   `json:"version"`
}

type Manifest struct {
	Skills      []Skill              `json:"skills"`
	Index       map[string]*Skill    `json:"-"`
	TrigIndex   map[string][]string  `json:"-"`
	DepIndex    map[string][]string  `json:"-"`
	GeneratedAt time.Time            `json:"generated_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

func NewManifest(skills []Skill) *Manifest {
	m := &Manifest{
		Skills:      skills,
		Index:       make(map[string]*Skill, len(skills)),
		TrigIndex:   make(map[string][]string),
		DepIndex:    make(map[string][]string),
		GeneratedAt: time.Now(),
		UpdatedAt:   time.Now(),
	}
	for i := range skills {
		sk := &skills[i]
		m.Index[sk.Slug] = sk
		for _, t := range sk.Triggers {
			m.TrigIndex[t] = append(m.TrigIndex[t], sk.Slug)
		}
		m.DepIndex[sk.Slug] = sk.Requires
	}
	return m
}

func (m *Manifest) Get(slug string) *Skill {
	return m.Index[slug]
}

func (m *Manifest) Count() int {
	return len(m.Skills)
}

func (m *Manifest) Search(query string) []*Skill {
	var result []*Skill
	for i := range m.Skills {
		sk := &m.Skills[i]
		if matchSkill(sk, query) {
			result = append(result, sk)
		}
	}
	return result
}

func matchSkill(sk *Skill, query string) bool {
	if contains(sk.Slug, query) || contains(sk.Name, query) || contains(sk.Description, query) {
		return true
	}
	for _, t := range sk.Triggers {
		if contains(t, query) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) && searchSubstring(s, substr)
}

func searchSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
