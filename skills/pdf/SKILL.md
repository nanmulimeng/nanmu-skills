---
name: pdf
description: 创建、读取、编辑、分析、处理 PDF 文件。当你涉及 PDF、读取 PDF、提取内容、表单填写、PDF 转换时必须使用此技能。即使用户只是说"读这个 PDF"或"把这个转成 PDF"，也应触发。如果未自动触发，手动写：`使用 pdf`。
---

# PDF Skill - PDF Document Processing

This skill enables reading, creating, editing, and processing PDF documents.

## Quick Reference

| Task | Tool |
|------|------|
| Read text content | `pypdf` or `pdfplumber` |
| Extract tables | `pdfplumber` |
| Create PDF | `reportlab` |
| Merge/split | `pypdf` or `qpdf` |
| OCR scanned PDFs | `pytesseract` + `pdf2image` |

## Python Libraries

### pypdf
- Merge, split, extract metadata
- Rotate pages
- Extract text (basic)

### pdfplumber
- Extract text with layout preservation
- Extract tables with accurate column detection
- Extract images

### reportlab
- Create PDFs from scratch
- Multi-page documents
- **Important**: For subscripts/superscripts, use XML markup tags (`<sub>`, `<sup>`) instead of Unicode characters

## Command-Line Tools

- `pdftotext` (poppler-utils) - Extract text from PDF
- `qpdf` - Merge, split, rotate, remove password
- `pdftk` - Merge, split, rotate (older)
- `pdfimages` - Extract images from PDF
- `pdfinfo` - Get PDF metadata

## Common Tasks

### OCR for Scanned PDFs
```bash
# Convert PDF to images
pdfimages -png input.pdf page

# OCR each image
pytesseract -l eng page-001.png
```

### Add Watermark
Use `pypdf` to merge watermark page with existing PDF

### Password Protection
Use `qpdf` to encrypt/decrypt PDFs

## Best Practices

1. Always verify text extraction accuracy
2. Check table extraction for alignment issues
3. Test PDF opens in standard viewers
4. Verify all pages are included
