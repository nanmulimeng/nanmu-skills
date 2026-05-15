---
name: coff0xc-skill-router
description: "Coff0xc 技能集索引与参考。列出全部 17 个 Coff0xc 综合技能的覆盖范围、来源和触发关键词。当用户明确询问 Coff0xc/coffee-skill 有哪些技能、某个 Coff0xc 技能覆盖什么内容、coff0xc 技能列表、Coff0xc 安全技能有哪些时使用。注意：全局路由请使用 skill-router；本索引仅用于查阅 Coff0xc 技能信息，不参与全局路由决策。中文触发：coff0xc 技能列表、coffee skill 有哪些、Coff0xc 包含什么、coff0xc 覆盖。英文触发：Coff0xc skill list、what coff0xc skills、coffee skill reference、list coff0xc capabilities。手动调用：使用 coff0xc-skill-router 查看技能列表。"
---

# coff0xc-skill-router — Coff0xc 技能集索引

## 目标

作为 Coff0xc 技能集的索引和参考，列出全部 17 个综合技能。**本文件是 Coff0xc 技能列表，不是全局路由器。** 全局跨域路由请使用 `skill-router`。

## 与 skill-router 的关系

| | skill-router | coff0xc-skill-router |
|---|---|---|
| **角色** | 全局唯一路由器 | Coff0xc 技能索引 |
| **覆盖** | 全部 55 个 skill | 17 个 Coff0xc skill |
| **触发** | "不确定用哪个" / "帮我分流" | "coff0xc 有哪些技能" |
| **功能** | 分析需求 → 推荐技能 | 列出 Coff0xc 技能详情 |

## 为什么需要索引

- Coff0xc 的 17 个技能是合并自 87 个源技能的综合技能，每个覆盖范围很广
- 用户可能不熟悉 Coff0xc 技能名和触发关键词
- 中文请求、缩写、英文工具名容易漏触发具体技能
- 当用户想了解 Coff0xc 有什么可用技能时，本索引用

## 路由规则

1. 如果另一个更具体的 Coff0xc skill 已经触发，优先执行具体 skill
2. 如果用户不确定用哪个 Coff0xc skill，建议先用 `skill-router` 分流，本索引提供 Coff0xc 技能详情
3. 如果请求涉及安全、生产、凭据、远程写入、删除或付费动作，套用目标 skill 的硬门禁
4. 如果无法确定路由，给出 2-3 个候选 Coff0xc skill 和最小澄清问题

## Coff0xc 技能列表

| 目标 skill | 合并来源 | 覆盖范围 |
| --- | --- | --- |
| coff0xc-software-engineering | c-cpp-dev, code-simplifier, git-workflow, go-dev, java-dev, js-ts-dev, python-dev, rust-dev, shell-scripting, testing | 全面软件工程：Python/JS/TS/Go/Rust/Java/C++/Shell、bugfix/feature、测试/重构、Git协作、构建质量 |
| coff0xc-ai-agent-rag | ai-agent-dev, ai-orchestrator, deep-thinking | AI Agent/RAG：Agent架构、工具调用、记忆系统、RAG管线、Prompt工程、多模型协作、评测/观测/成本 |
| coff0xc-api-data-platform | api-design, database, cli-creator | API/数据平台：REST/GraphQL/OpenAPI、数据库schema/迁移、CLI/SDK、分页/认证/错误码、ETL/数据质量 |
| coff0xc-ui-doc-output | UIdesign, pdf, quick-translate | UI/文档输出：产品UI设计系统、交互状态、PDF/Word/PPT文档处理、报告交付、翻译润色 |
| coff0xc-research-drawio-diagram | drawio-generator | 科研绘图：draw.io/diagrams.net、论文配图、算法架构图、模型结构图、可编辑.drawio生成 |
| coff0xc-secure-code-appsec | api-discovery, api-security-test, backdoor-detector, browser-security, code-audit, graphql-pentest, llm-red-teaming, oauth-security, spa-pentest, web-pentest | 代码/应用安全：代码审计、source/sink污点分析、Web/API/GraphQL/OAuth安全、LLM安全、后门检测、SPA安全 |
| coff0xc-cloud-devsecops | cloud-security, container-security, devsecops, docker-k8s, secrets-management, serverless-security, supply-chain-security | 云/DevSecOps：AWS/Azure/GCP、Docker/K8s、Serverless、CI/CD、供应链安全、IaC/Terraform、密钥管理 |
| coff0xc-detection-response | detection-engineering, email-security, forensics-analysis, incident-response, malware-analysis, osint, soc-operations, threat-hunting, threat-intelligence | 检测/响应：SIEM/Sigma/YARA、SOC运营、威胁狩猎/情报、应急响应、取证、恶意软件分析、邮件安全 |
| coff0xc-vulnerability-lifecycle | bug-bounty, pentest-report, red-team-poc, vuln-research, vulnerability-management | 漏洞生命周期：CVE研究、补丁分析、CVSS/EPSS/KEV优先级、PoC验证、漏洞报告、修复跟踪 |
| coff0xc-identity-zero-trust | ad-pentest, credential-access, identity-security, lateral-movement, privilege-escalation, zero-trust | 身份/零信任：IAM/SSO/MFA、AD/Kerberos/BloodHound、凭证风险、提权/横向移动防御、Zero Trust/PAM |
| coff0xc-authorized-assessment | attack-chain-orchestrator, autoredteam-orchestrator, c2-framework, cdn-bypass, data-exfiltration, evasion-toolkit, fingerprint-engine, full-pentest, phishing-simulation, post-exploitation, proxy-pool-manager, recon-workflow, red-team-infra, security-tool-dev, social-engineering | 授权评估：ROE/授权边界、攻击面梳理、红队防御化、C2/evasion、钓鱼/外传演练、安全工具开发 |
| coff0xc-binary-mobile-iot | binary-exploit, crypto-security, ctf, ics-scada, iot-security, kernel-security, mobile-security, reverse-engineering | 二进制/移动/IoT：逆向工程、PWN/CTF、内核安全、Frida/APK/IPA、固件分析、IoT/ICS/SCADA、密码学审计 |
| coff0xc-blockchain-security | blockchain-security | 区块链安全：Solidity/EVM、Solana/Cosmos/Substrate、Cairo/StarkNet、TON/Algorand、DeFi/AMM/oracle、Token/NFT、Foundry/Hardhat/Slither |
| coff0xc-compliance-architecture | compliance-audit, data-security, security-architecture | 合规/架构：STRIDE威胁建模、等保/PCI-DSS/GDPR/ISO27001/SOC2、CIS/NIST基线、数据安全/DLP/隐私、成熟度评估 |
| coff0xc-purple-deception | honeypot, purple-team | 紫队/欺骗：ATT&CK映射、紫队演练计划、检测/响应验证、蜜罐/deception/canary、覆盖指标 |
| coff0xc-network-protocol-security | network-protocol, wireless-security | 网络协议安全：TLS/DNS/HTTP3/QUIC、TCP/UDP分析、pcap/Wireshark、无线/BLE/RF、ProVerif形式化建模 |

## 手动触发写法

如果自动触发不稳定，用户可以直接写：

```text
使用 coff0xc-secure-code-appsec 审计这个项目
使用 coff0xc-ai-agent-rag 设计一个 RAG Agent
使用 coff0xc-cloud-devsecops 检查 K8s 和 CI/CD
使用 coff0xc-detection-response 写 Sigma/YARA 检测
使用 coff0xc-research-drawio-diagram 画算法架构图
使用 coff0xc-skill-router 查看 Coff0xc 技能列表
```

## 触发排障

- skill 目录名必须和 frontmatter `name` 一致
- 安装或替换后需要重启/刷新客户端，让 skill 列表重新索引
- frontmatter 只保留 `name` 和 `description` 最稳；触发词必须写进 `description`
- 同名 skill 分布在多个目录时可能抢占触发，保留一份主版本
- 太短或只写抽象能力的 description 容易漏触发；应包含中文、英文、工具名、任务名和常见缩写
