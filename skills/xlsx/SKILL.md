---
name: xlsx
description: 创建、读取、编辑、分析 Excel 电子表格（.xlsx 文件）。当你涉及 Excel、电子表格、创建表格、读取 .xlsx、数据分析时必须使用此技能。即使用户只是说"处理这个 xlsx"或"生成一份 Excel 报表"，也应触发。如果未自动触发，手动写：`使用 xlsx`。
---

# XLSX Skill - Excel Spreadsheet Processing

This skill handles any .xlsx file operations including creation, reading, editing, and data analysis.

## Quick Reference

| Task | Approach |
|------|----------|
| Read/analyze content | `pandas` for data, `openpyxl` for formulas |
| Create new spreadsheet | Use `openpyxl` or generate with Python |
| Edit existing spreadsheet | Open with openpyxl, modify, save |

## Output Requirements

- **Professional font**: Use Calibri or Arial (Excel defaults)
- **Zero formula errors**: Always verify formulas work
- **Preserve existing templates**: Don't break user's formatting

## Financial Model Standards

### Color Coding
- **Blue**: Input values (hardcoded numbers)
- **Black**: Formulas
- **Green**: Internal links/references
- **Red**: External links/references

### Number Formatting
- Years as text (e.g., "2024")
- Currency: `$#,##0`
- Zeros display as "-"
- Percentages: one decimal place

## Formula Rules

1. **Use cell references** - Never hardcode calculated values
2. **Document assumptions** - Add comments for complex formulas
3. **Verify before delivery** - Open file and check for #REF!, #DIV/0!, etc.

## Common Tools

- `pandas` - Data analysis and manipulation
- `openpyxl` - Formula creation and formatting
- `LibreOffice` - Formula recalculation
- `xlsxwriter` - Creating new files with formatting

## Verification Checklist

Before delivering any xlsx file:
- [ ] Open and verify all formulas calculate correctly
- [ ] Check for #REF!, #DIV/0!, #VALUE! errors
- [ ] Verify number formatting matches requirements
- [ ] Test all sheets are accessible
