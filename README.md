# nanmu-skills — 泰坦机魂技能集

Claude Code 技能库，涵盖全栈开发、安全审计、移动开发、飞书生态、文档处理等 55 个技能。

## 目录结构

```
nanmu-skills/
├── mcp-skills/          ← 32 个非飞书技能 → 部署到自建 MCP 服务
│   ├── coff0xc-software-engineering/   (综合，覆盖 Go/Python/JS/Shell/Git/测试)
│   ├── coff0xc-api-data-platform/      (综合，覆盖 API/数据库)
│   ├── coff0xc-secure-code-appsec/     (综合，覆盖代码审计)
│   ├── coff0xc-binary-mobile-iot/      (综合，覆盖移动安全/逆向)
│   ├── ... 17 个 coff0xc 综合技能
│   ├── backend-engineering/            (独特：架构/部署/进程管理)
│   ├── perf-engineering/               (独特：性能优化)
│   └── ... 15 个独特原子技能
├── lark-skills/         ← 23 个飞书技能 → 同步到本地 ~/.claude/skills/
│   ├── lark-approval/
│   ├── lark-base/
│   └── ...
└── evals/               ← 评测用例
```

## 技能体系

| 分类 | 目录 | 数量 | 部署位置 |
|------|------|------|----------|
| 非飞书技能 | `mcp-skills/` | 32 | 自建 MCP 服务器 (123.56.223.97:3456) |
| 飞书技能 | `lark-skills/` | 23 | 本地 `~/.claude/skills/` |

### mcp-skills (32)

| 领域 | 技能 |
|------|------|
| coff0xc 综合技能 (17) | coff0xc-software-engineering, coff0xc-api-data-platform, coff0xc-secure-code-appsec, coff0xc-binary-mobile-iot, coff0xc-ai-agent-rag, coff0xc-authorized-assessment, coff0xc-blockchain-security, coff0xc-cloud-devsecops, coff0xc-compliance-architecture, coff0xc-detection-response, coff0xc-identity-zero-trust, coff0xc-network-protocol-security, coff0xc-purple-deception, coff0xc-research-drawio-diagram, coff0xc-skill-router, coff0xc-ui-doc-output, coff0xc-vulnerability-lifecycle |
| dev-engineering (2) | backend-engineering, perf-engineering, web-access |
| apple-mobile (3) | apple-development, flutter-development, uniapp-dev |
| documents (4) | docx, xlsx, pdf, pptx |
| design-product (2) | ui-design, product-manager |
| other (2) | interview-prep, pua |
| meta (1) | skill-router |

> 注：原 11 个原子开发/安全技能 (go-dev, python-dev, js-ts-dev, shell-scripting, git-workflow, test-engineering, api-design, db-design, code-audit, mobile-security, reverse-engineering) 已被对应的 coff0xc 综合技能完全覆盖，不再单独维护。

### lark-skills (23)

飞书生态全覆盖：审批、考勤、多维表格、日历、通讯录、文档、云盘、事件、IM、邮箱、妙记、OKR、OpenAPI、共享、电子表格、Skill Maker、演示文稿、任务、视频会议、白板、知识库、会议纪要、日报。

## 部署方式

```bash
# MCP 服务更新
cd mcp-skills/ && tar czf deploy.tar.gz * && scp deploy.tar.gz nanmu@123.56.223.97:/tmp/
ssh nanmu@123.56.223.97 "cd /home/nanmu/nanmu-skill-mcp/skills && rm -rf * && tar xzf /tmp/deploy.tar.gz && sudo systemctl restart nanmu-skill-mcp"

# 本地飞书技能同步
cp -r lark-skills/* ~/.claude/skills/
```

## 许可

MIT License
