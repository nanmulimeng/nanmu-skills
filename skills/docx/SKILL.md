---
name: docx
description: Create, read, edit, and analyze Word documents (.docx files). Use when user mentions docx, Word document, or needs to work with .docx files.
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
python scripts/office/unpack.py document.docx unpacked/
```

**Step 2: Edit XML**
- Use "Claude" as author for tracked changes
- Use Edit tool directly for string replacement
- Use smart quotes (XML entities like `&#x2019;`)

**Step 3: Pack**
```bash
python scripts/office/pack.py unpacked/ output.docx --original document.docx
```

## Critical Rules Summary

- Always use `WidthType.DXA` for tables (not PERCENTAGE)
- Include `outlineLevel` in headings for TOC
- Use `ShadingType.CLEAR` for table shading
- Never use tables as dividers - use paragraph borders instead
- Override built-in styles with exact IDs ("Heading1", "Heading2", etc.)
