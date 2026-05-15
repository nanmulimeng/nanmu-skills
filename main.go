package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"nanmu-skill-mcp/mcp"
	"nanmu-skill-mcp/registry"
	"nanmu-skill-mcp/tools"
)

func main() {
	addr := flag.String("addr", ":3456", "HTTP listen address")
	skillsDir := flag.String("skills-dir", "skills", "技能文件目录")
	authKey := flag.String("key", "", "认证 key (X-Skill-Key header)")
	flag.Parse()

	if *authKey == "" {
		*authKey = os.Getenv("SKILL_KEY")
	}

	reg, err := registry.LoadFromDir(*skillsDir)
	if err != nil {
		log.Fatalf("加载技能失败: %v", err)
	}
	log.Printf("已加载 %d 个技能", reg.Count())

	srv := mcp.NewServer(*authKey)
	rt := &tools.RouterTool{Manifest: reg}

	srv.RegisterTool("route_skill",
		"根据需求和项目证据路由 must/conditional/skip 技能",
		map[string]mcp.Property{
			"request":      {Type: "string", Description: "用户的自然语言需求"},
			"project_path": {Type: "string", Description: "项目路径 (可选)"},
			"mode":         {Type: "string", Description: "模式 (可选)"},
		},
		[]string{"request"},
		rt.RouteSkill,
	)

	srv.RegisterTool("fetch_skill_raw",
		"拉取指定技能 raw 或摘要",
		map[string]mcp.Property{
			"slug":   {Type: "string", Description: "单个 skill slug"},
			"slugs":  {Type: "array", Description: "多个 skill slugs", Items: &mcp.Items{Type: "string"}},
			"format": {Type: "string", Description: "raw 或 summary"},
		},
		nil,
		rt.FetchSkillRaw,
	)

	srv.RegisterTool("explain_route",
		"解释路由结果",
		map[string]mcp.Property{
			"request":      {Type: "string", Description: "用户的自然语言需求"},
			"project_path": {Type: "string", Description: "项目路径 (可选)"},
		},
		[]string{"request"},
		rt.ExplainRoute,
	)

	srv.RegisterTool("search_manifest",
		"在 manifest 中搜索技能候选",
		map[string]mcp.Property{
			"query": {Type: "string", Description: "搜索关键词"},
		},
		[]string{"query"},
		rt.SearchManifest,
	)

	srv.RegisterTool("validate_skill",
		"校验任务是否遗漏必须技能",
		map[string]mcp.Property{
			"task":     {Type: "string", Description: "任务描述"},
			"selected": {Type: "array", Description: "已选技能列表", Items: &mcp.Items{Type: "string"}},
		},
		[]string{"task", "selected"},
		rt.ValidateSkill,
	)

	srv.RegisterTool("doctor",
		"检查配置、网络、缓存和 key",
		nil, nil,
		rt.Doctor,
	)

	srv.RegisterTool("refresh_cache",
		"刷新 manifest 缓存",
		nil, nil,
		rt.RefreshCache,
	)

	srv.RegisterTool("get_cache_status",
		"查看缓存状态",
		nil, nil,
		rt.GetCacheStatus,
	)

	go func() {
		log.Printf("nanmu-skill-mcp 启动于 %s (skills: %d)", *addr, reg.Count())
		if err := srv.ListenAndServe(*addr); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("正在关闭...")
}
