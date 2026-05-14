# nanmu-skills — 泰坦机魂技能集

Claude Code 技能市场，涵盖全栈开发、安全审计、移动开发、飞书生态、文档处理等 48 个技能。

## 安装

```bash
# 注册市场
/plugin marketplace add nanmulimeng/nanmu-skills

# 安装全部技能
/plugin install dev-engineering@nanmu-skills
/plugin install security@nanmu-skills
/plugin install apple-mobile@nanmu-skills
/plugin install lark@nanmu-skills
/plugin install documents@nanmu-skills
/plugin install design-product@nanmu-skills
/plugin install other@nanmu-skills
```

## 插件列表

| 插件 | 技能数 | 说明 |
|------|--------|------|
| `dev-engineering` | 11 | API设计、后端、数据库、Git、Go/Python/JS/Shell、测试、性能、Web |
| `security` | 3 | 代码审计、移动安全、逆向工程 |
| `apple-mobile` | 3 | Apple原生、Flutter、uni-app |
| `lark` | 23 | 飞书生态全覆盖 |
| `documents` | 4 | Word/Excel/PDF/PPT |
| `design-product` | 2 | UI设计、产品经理 |
| `other` | 2 | 面试准备、PUA识别 |

## 技能结构

每个技能目录包含一个 `SKILL.md` 文件，部分含 `references/` 参考文档：

```
skills/
  <skill-name>/
    SKILL.md       # 技能定义（YAML frontmatter + 指令）
    references/    # 可选参考文档
```

## 许可

MIT License
