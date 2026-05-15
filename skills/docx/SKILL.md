---
name: docx
description: 创建、读取、编辑、分析 Word 文档（.docx 文件）。当你涉及 Word 文档、创建文档、编辑文档、读取 .docx 文件时必须使用此技能。即使用户只是说"处理这个 docx"或"生成一份 Word 报告"，也应触发。如果未自动触发，手动写：`使用 docx`。
---

# DOCX Skill - Word Document Processing

This skill enables creation, reading, editing, and manipulation of Word documents (.docx files). A .docx file is a ZIP archive containing XML files.

## Quick Reference

| Task | Approach |
|------|----------|
| Read/analyze content | `pandoc` or unpack for raw XML |
| Create new document | Use `docx-js` library |
| Edit existing document | Unpack → edit XML → repack |

## Creating New Documents

Use the `docx` JavaScript library to generate .docx files. Key requirements:

- **Set page size explicitly** - docx-js defaults to A4, not US Letter (US Letter: 12240 x 15840 DXA)
- **Use landscape orientation** - pass portrait dimensions and let docx-js swap internally
- **Never use `\n`** - use separate Paragraph elements
- **Never use unicode bullets** - use `LevelFormat.BULLET` with numbering config
- **Table dual widths** - set both `columnWidths` on table AND `width` on each cell using DXA
- **ImageRun requires `type`** - always specify png/jpg/etc
- **PageBreak must be in Paragraph** - standalone creates invalid XML

## Editing Existing Documents

**Step 1: Unpack**
```bash
# .docx is a ZIP archive; unpack with any ZIP tool
mkdir -p unpacked && cd unpacked && unzip ../document.docx
```

**Step 2: Edit XML**
- Use "Claude" as author for tracked changes
- Use Edit tool directly for string replacement
- Use smart quotes (XML entities like `&#x2019;`)

**Step 3: Pack**
```bash
# Repack the modified XML back into a .docx (ZIP archive)
cd unpacked && zip -r ../output.docx . -x "*.DS_Store"
```

## Critical Rules Summary

- Always use `WidthType.DXA` for tables (not PERCENTAGE)
- Include `outlineLevel` in headings for TOC
- Use `ShadingType.CLEAR` for table shading
- Never use tables as dividers - use paragraph borders instead
- Override built-in styles with exact IDs ("Heading1", "Heading2", etc.)
