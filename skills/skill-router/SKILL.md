---
name: skill-router
description: "全局技能路由器——当用户需求跨越多个技能域、不确定该用哪个技能、或特定技能未自动触发时的兜底分流。覆盖全部 55 个技能域：独立技能(后端/性能/Web/Apple/Flutter/uni-app/飞书23个/文档4个/设计2个/其他2个) + Coff0xc综合技能(17个)。中文触发：不确定用哪个 skill、帮我选择技能、同时涉及多个领域、这个需求该用哪个技能、skill 路由、分流、跨领域、混合需求。英文触发：which skill、skill router、what skill should I use、how should I handle this、not sure which skill。手动调用：使用 skill-router 帮我分流。"
---

# 全局技能路由器 (Universal Skill Router)

## 目标

当用户需求不明确、跨多个技能域、或没有特定技能自动命中时，帮用户判断应该使用哪个（或哪几个）技能。**本 router 是唯一全局路由入口，覆盖全部 55 个 skills（独立技能 38 + Coff0xc 综合 17）。**

## 技能体系

本技能库包含两套互补的技能体系：

| 体系 | 来源 | 数量 | 粒度 | 适用场景 |
|------|------|------|------|----------|
| **独立技能** | nanmu-skills | 38 | 专题/领域 | 单一领域明确任务（如"部署后端"、"查飞书审批"、"生成PPT"） |
| **Coff0xc 综合** | coffee-skill | 17 | 综合/覆盖广 | 跨域复杂任务（如"设计一个带安全审计的全栈系统"） |

**路由原则**：单一明确领域 → 独立技能；跨域/复杂/不确定 → Coff0xc 综合技能；仍不确定 → 本 router 分流。

## 适用场景

- 用户说"不确定该用哪个技能"或"帮我选择"
- 需求同时涉及多个领域（如"写个带审计日志的API"=API平台+代码安全）
- 用户用口语描述需求，没有命中任何技能的触发词
- 用户说"这个有点复杂"或"涉及面比较多"
- 用户问"用哪个 coff0xc skill"或提到 Coff0xc 技能

## 不适用场景

- 用户明确说了某个技能覆盖范围内的需求（直接命中即可）
- 简单的信息问答（"xxx是什么"）

## 路由决策流程

### 第一步：识别需求涉及的能力域

根据用户描述，提取关键词，映射到以下能力域：

#### 独立技能 (nanmu-skills)

| 能力域 | 关键词 | 对应插件 | 核心技能 |
|--------|--------|----------|----------|
| 后端架构 | 服务端、中间件、架构、部署、配置 | dev-engineering | backend-engineering |
| 性能 | 优化、慢查询、内存泄漏、N+1、缓存 | dev-engineering | perf-engineering |
| Web可访问性 | 网页、浏览器、爬虫、CDP、截图、Playwright | dev-engineering | web-access |
| Apple | Swift、SwiftUI、Xcode、iOS、macOS | apple-mobile | apple-development |
| Flutter | Flutter、Dart、跨平台、Widget | apple-mobile | flutter-development |
| uni-app | uni-app、小程序、微信、H5、跨端 | apple-mobile | uniapp-dev |
| 飞书 | 飞书、lark、审批、考勤、IM、文档等 | lark | lark-* (23个) |
| Word | Word、docx、报告、排版、封面 | documents | docx |
| Excel | Excel、xlsx、透视表、公式、图表 | documents | xlsx |
| PDF | PDF、提取、合并、表单、签名 | documents | pdf |
| PPT | PPT、pptx、演示、幻灯片、动画 | documents | pptx |
| UI设计 | 界面、样式、布局、颜色、组件、设计 | design-product | ui-design |
| 产品 | 需求、功能、用户价值、竞品分析、PRD | design-product | product-manager |
| 面试 | 面试、八股文、Java基础、JVM、MySQL | other | interview-prep |
| 高压模式 | 加油、PUA、别偷懒、为什么还不行 | other | pua |

> 原子开发技能（go-dev/python-dev/js-ts-dev/shell-scripting/git-workflow/test-engineering/api-design/db-design）已合并入 coff0xc-software-engineering；安全原子技能（code-audit/mobile-security/reverse-engineering）已合并入对应 coff0xc 综合技能。

#### Coff0xc 综合技能 (coffee-skill)

| 能力域 | 关键词 | 核心技能 |
|--------|--------|----------|
| 综合软件工程 | 多语言项目、bugfix+test+git全流程、代码重构规模化、Go/Python/JS/Shell开发 | coff0xc-software-engineering |
| AI Agent/RAG | Agent架构、RAG管线、Prompt工程、多模型协作、LLM评测、向量数据库、embedding | coff0xc-ai-agent-rag |
| API数据平台 | REST/GraphQL设计、OpenAPI、数据库schema、CLI/SDK、数据模型、ETL、兼容演进 | coff0xc-api-data-platform |
| UI文档输出 | 产品UI设计系统、PDF/Word/PPT文档、报告交付、技术翻译润色 | coff0xc-ui-doc-output |
| 科研绘图 | draw.io、diagrams.net、论文配图、算法架构图、模型结构图、Neural Network | coff0xc-research-drawio-diagram |
| 代码应用安全 | 代码审计、source/sink、OAuth/GraphQL安全、LLM安全、后门检测、SPA安全 | coff0xc-secure-code-appsec |
| 云DevSecOps | AWS/Azure/GCP、Docker/K8s、CI/CD、供应链安全、IaC/Terraform、密钥管理 | coff0xc-cloud-devsecops |
| 检测响应 | SIEM/Sigma/YARA、SOC运营、威胁狩猎、应急响应、取证、恶意软件分析 | coff0xc-detection-response |
| 漏洞生命周期 | CVE研究、补丁分析、CVSS/EPSS/KEV、漏洞管理、PoC验证、修复跟踪 | coff0xc-vulnerability-lifecycle |
| 身份零信任 | IAM/SSO/MFA、AD/Kerberos、BloodHound、提权/横向移动、Zero Trust/PAM | coff0xc-identity-zero-trust |
| 授权评估 | 红队演练、攻击链、C2/evasion、钓鱼演练、数据外传、ROE/授权边界 | coff0xc-authorized-assessment |
| 二进制移动IoT | 逆向工程、PWN/CTF、Frida/APK/IPA、固件/IoT/ICS、内核安全、密码学审计 | coff0xc-binary-mobile-iot |
| 区块链安全 | Solidity/EVM、Solana/Cosmos/Substrate、DeFi/NFT/Token、Foundry/Hardhat/Slither | coff0xc-blockchain-security |
| 合规架构 | STRIDE威胁建模、等保/ISO27001/GDPR/SOC2、DLP/隐私、安全基线/成熟度 | coff0xc-compliance-architecture |
| 紫队欺骗 | ATT&CK映射、紫队演练、检测验证、蜜罐/deception/canary、覆盖指标 | coff0xc-purple-deception |
| 网络协议安全 | TLS/DNS/HTTP3/QUIC、pcap/Wireshark、无线/BLE/RF、ProVerif形式化 | coff0xc-network-protocol-security |

### 第二步：选择技能粒度

```
需求明确单一？
  ├── 是 → 使用独立技能（如 backend-engineering, docx, apple-development）
  └── 否 → 涉及多域/复杂流程？
            ├── 是 → 使用 Coff0xc 综合技能（如 coff0xc-software-engineering）
            └── 不确定 → 本 router 给出 2-3 个候选
```

### 第三步：输出路由建议

```
## 技能路由分析

**需求拆解：**
- [领域A]：涉及xxx → 推荐 `skill-a`（独立）/ `coff0xc-xxx`（综合）
- [领域B]：涉及xxx → 推荐 `skill-b`

**推荐操作：**
如果主要关注 [最领域]，可先说 "使用 [技能名]" 手动触发。
如需多技能配合，我会依次加载处理。
```

### 第四步：按需加载

路由完成后，按用户选择加载对应技能内容执行。

## 触发强化

- 自动触发依赖本文件 frontmatter 的 `description`
- 如果没有自动触发，手动写：`使用 skill-router 帮我分流`
- 本 skill 覆盖了全部 55 个技能域的关键词以最大化命中概率

## 跨域常见组合

| 用户典型场景 | 推荐技能组合 |
|-------------|-------------|
| 写代码+提交+测试 | coff0xc-software-engineering |
| 开发+性能优化 | coff0xc-software-engineering + perf-engineering |
| 开发+安全检查 | coff0xc-software-engineering + coff0xc-secure-code-appsec |
| API设计+数据库+文档 | coff0xc-api-data-platform + docx/pdf |
| UI+前端逻辑 | ui-design + flutter-development/uniapp-dev |
| 飞书操作+文档 | lark-* + docx/xlsx |
| 全栈项目(多语言+测试+Git) | coff0xc-software-engineering |
| AI系统设计(Agent+RAG+评测) | coff0xc-ai-agent-rag |
| 安全审计+云架构+合规 | coff0xc-secure-code-appsec + coff0xc-cloud-devsecops |
| 移动安全+逆向+二进制 | coff0xc-binary-mobile-iot |
| 红队演练+检测验证 | coff0xc-authorized-assessment + coff0xc-purple-deception |

## Coff0xc 技能手动触发

```text
使用 coff0xc-software-engineering 重构这个多语言项目
使用 coff0xc-ai-agent-rag 设计一个 RAG Agent
使用 coff0xc-secure-code-appsec 审计代码安全
使用 coff0xc-cloud-devsecops 检查 K8s 和 CI/CD
使用 coff0xc-detection-response 写 Sigma/YARA 检测
使用 coff0xc-skill-router 查看 Coff0xc 技能列表
```
