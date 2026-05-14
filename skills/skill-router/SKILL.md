---
name: skill-router
description: "泰坦机魂技能路由器——当用户需求跨越多个技能域、不确定该用哪个技能、或特定技能未自动触发时的兜底分流。覆盖全部48个技能域：开发工程(API设计/后端/数据库/Git/Go/Python/JS/Shell/测试/性能/Web)、安全(代码审计/移动安全/逆向)、Apple移动(Apple原生/Swift/SwiftUI/Flutter/uni-app/小程序)、飞书生态(审批/考勤/多维表格/日历/通讯录/文档/云盘/IM/邮箱/妙记/OKR/电子表格/演示文稿/任务/视频会议/白板/知识库)、文档(docx/xlsx/pdf/pptx)、设计产品(UI设计/产品经理)、其他(面试准备/八股文/PUA模式)。中文触发：不确定用哪个 skill、帮我选择技能、同时涉及多个领域、这个需求该用哪个技能、skill 路由、分流、跨领域、混合需求。英文触发：which skill、skill router、what skill should I use、how should I handle this、not sure which skill。手动调用：使用 skill-router 帮我分流。"
---

# 泰坦机魂技能路由器 (Skill Router)

## 目标

当用户需求不明确、跨多个技能域、或没有特定技能自动命中时，帮用户判断应该使用哪个（或哪几个）技能。

## 适用场景

- 用户说"不确定该用哪个技能"或"帮我选择"
- 需求同时涉及多个领域（如"写个带审计日志的API"=API设计+代码审计）
- 用户用口语描述需求，没有命中任何技能的触发词
- 用户说"这个有点复杂"或"涉及面比较多"

## 不适用场景

- 用户明确说了某个技能覆盖范围内的需求（直接命中即可）
- 简单的信息问答（"xxx是什么"）

## 路由决策流程

### 第一步：识别需求涉及的能力域

根据用户描述，提取关键词，映射到以下能力域：

| 能力域 | 关键词 | 对应插件 | 核心技能 |
|--------|--------|----------|----------|
| API设计 | 接口、端点、REST、路由、HTTP、响应格式 | dev-engineering | api-design |
| 后端架构 | 服务端、中间件、架构、部署、配置 | dev-engineering | backend-engineering |
| 数据库 | SQL、表结构、索引、查询、字段、迁移 | dev-engineering | db-design |
| Git | 提交、commit、分支、合并、push、PR | dev-engineering | git-workflow |
| Go | Golang、goroutine、channel、go mod | dev-engineering | go-dev |
| JS/TS | JavaScript、TypeScript、Node、npm、ESM | dev-engineering | js-ts-dev |
| Python | Python、pip、FastAPI、Django、pytest | dev-engineering | python-dev |
| Shell | Bash、脚本、shell、自动化、部署脚本 | dev-engineering | shell-scripting |
| 测试 | 测试、TDD、覆盖率、Mock、断言、单测 | dev-engineering | test-engineering |
| 性能 | 优化、慢查询、内存泄漏、N+1、缓存 | dev-engineering | perf-engineering |
| Web可访问性 | 网页、浏览器、爬虫、CDP、截图 | dev-engineering | web-access |
| 代码审计 | 安全审计、漏洞、注入、XSS、OWASP | security | code-audit |
| 移动安全 | APK、IPA、Frida、Hook、SSL Pinning | security | mobile-security |
| 逆向 | 逆向、反编译、IDA、Ghidra、脱壳 | security | reverse-engineering |
| Apple | Swift、SwiftUI、Xcode、iOS、macOS | apple-mobile | apple-development |
| Flutter | Flutter、Dart、跨平台、Widget | apple-mobile | flutter-development |
| uni-app | uni-app、小程序、微信、H5、跨端 | apple-mobile | uniapp-dev |
| 飞书 | 飞书、lark、审批、考勤、IM、文档等 | lark | lark-* (23个) |
| 文档 | Word、Excel、PDF、PPT、报告 | documents | docx/xlsx/pdf/pptx |
| UI设计 | 界面、样式、布局、颜色、组件、设计 | design-product | ui-design |
| 产品 | 需求、功能、用户价值、竞品分析 | design-product | product-manager |
| 面试 | 面试、八股文、Java基础、JVM、MySQL | other | interview-prep |
| 高压模式 | 加油、PUA、别偷懒、为什么还不行 | other | pua |

### 第二步：输出路由建议

```
## 技能路由分析

**需求拆解：**
- [领域A]：涉及xxx → 推荐 `skill-a`
- [领域B]：涉及xxx → 推荐 `skill-b`

**推荐操作：**
如果你主要关注 [最主要领域]，可以先说 "使用 [技能名]" 来手动触发。
如果确实需要多个技能配合，我会依次加载处理。
```

### 第三步：按需加载

路由完成后，按用户选择加载对应技能内容执行。

## 触发强化

- 自动触发依赖本文件 frontmatter 的 `description`
- 如果没有自动触发，手动写：`使用 skill-router 帮我分流`
- 本 skill 覆盖了全部 48 个技能域的关键词以最大化命中概率

## 跨域常见组合

| 用户典型场景 | 推荐技能组合 |
|-------------|-------------|
| 写一个新API+数据库表 | api-design + db-design |
| 写代码+提交 | go-dev/js-ts-dev/python-dev + git-workflow |
| 开发+性能优化 | go-dev/js-ts-dev/python-dev + perf-engineering |
| 开发+安全检查 | 语言技能 + code-audit |
| API+文档 | api-design + docx/pdf |
| UI+前端逻辑 | ui-design + js-ts-dev/flutter-development/uniapp-dev |
| 飞书操作+文档 | lark-* + docx/xlsx |
