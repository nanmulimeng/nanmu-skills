package tools

import (
	"encoding/json"

	"nanmu-skill-mcp/registry"
	"nanmu-skill-mcp/router"
)

type RouterTool struct {
	Manifest *registry.Manifest
}

func (t *RouterTool) RouteSkill(args json.RawMessage) (interface{}, error) {
	var params struct {
		Request     string `json:"request"`
		ProjectPath string `json:"project_path,omitempty"`
		Mode        string `json:"mode,omitempty"`
	}
	if err := json.Unmarshal(args, &params); err != nil {
		return nil, err
	}
	result := router.Route(t.Manifest, params.Request)
	return result, nil
}

func (t *RouterTool) FetchSkillRaw(args json.RawMessage) (interface{}, error) {
	var params struct {
		Slug   string   `json:"slug,omitempty"`
		Slugs  []string `json:"slugs,omitempty"`
		Format string   `json:"format,omitempty"`
	}
	if err := json.Unmarshal(args, &params); err != nil {
		return nil, err
	}

	slugs := params.Slugs
	if params.Slug != "" {
		slugs = append(slugs, params.Slug)
	}

	type SkillResult struct {
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Body        string `json:"body,omitempty"`
	}

	var results []SkillResult
	for _, slug := range slugs {
		sk := t.Manifest.Get(slug)
		if sk == nil {
			continue
		}
		sr := SkillResult{
			Slug:        sk.Slug,
			Name:        sk.Name,
			Description: sk.Description,
		}
		if params.Format != "summary" {
			sr.Body = sk.Body
		}
		results = append(results, sr)
	}

	return map[string]interface{}{"skills": results}, nil
}

func (t *RouterTool) ExplainRoute(args json.RawMessage) (interface{}, error) {
	var params struct {
		Request     string `json:"request"`
		ProjectPath string `json:"project_path,omitempty"`
	}
	if err := json.Unmarshal(args, &params); err != nil {
		return nil, err
	}

	result := router.Route(t.Manifest, params.Request)

	type ExplainItem struct {
		Slug    string   `json:"slug"`
		Name    string   `json:"name"`
		Reason  string   `json:"reason"`
		Matches []string `json:"matches,omitempty"`
	}

	var items []ExplainItem
	for _, slug := range result.Must {
		sk := t.Manifest.Get(slug)
		if sk != nil {
			items = append(items, ExplainItem{
				Slug:   slug,
				Name:   sk.Name,
				Reason: "高匹配度 + 触发词命中",
			})
		}
	}
	for _, slug := range result.Conditional {
		sk := t.Manifest.Get(slug)
		if sk != nil {
			items = append(items, ExplainItem{
				Slug:   slug,
				Name:   sk.Name,
				Reason: "中等匹配度，视具体情况加载",
			})
		}
	}

	return map[string]interface{}{"explanations": items, "route": result}, nil
}

func (t *RouterTool) SearchManifest(args json.RawMessage) (interface{}, error) {
	var params struct {
		Query string `json:"query"`
	}
	if err := json.Unmarshal(args, &params); err != nil {
		return nil, err
	}

	results := t.Manifest.Search(params.Query)
	type SearchResult struct {
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var items []SearchResult
	for _, sk := range results {
		items = append(items, SearchResult{
			Slug:        sk.Slug,
			Name:        sk.Name,
			Description: sk.Description,
		})
	}
	return map[string]interface{}{"results": items, "count": len(items)}, nil
}

func (t *RouterTool) ValidateSkill(args json.RawMessage) (interface{}, error) {
	var params struct {
		Selected []string `json:"selected"`
		Task     string   `json:"task"`
	}
	if err := json.Unmarshal(args, &params); err != nil {
		return nil, err
	}

	result := router.Route(t.Manifest, params.Task)

	selected := make(map[string]bool)
	for _, s := range params.Selected {
		selected[s] = true
	}

	var missing []string
	for _, s := range result.Must {
		if !selected[s] {
			missing = append(missing, s)
		}
	}

	return map[string]interface{}{"missing": missing, "task": params.Task}, nil
}

func (t *RouterTool) Doctor(args json.RawMessage) (interface{}, error) {
	return map[string]interface{}{
		"status":        "healthy",
		"skills_count":  t.Manifest.Count(),
		"config":        "ok",
		"cache":         "memory",
		"network":       "n/a (self-hosted)",
		"version":       "1.0.0",
	}, nil
}

func (t *RouterTool) RefreshCache(args json.RawMessage) (interface{}, error) {
	return map[string]interface{}{
		"refreshed":    true,
		"skills_count": t.Manifest.Count(),
	}, nil
}

func (t *RouterTool) GetCacheStatus(args json.RawMessage) (interface{}, error) {
	return map[string]interface{}{
		"skills_count": t.Manifest.Count(),
		"updated_at":   t.Manifest.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
