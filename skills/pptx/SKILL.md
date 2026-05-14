---
name: pptx
description: Create, read, edit, analyze PowerPoint presentations (.pptx files). Use when user mentions pptx, PowerPoint, slides, presentation, or deck.
---

# PPTX Skill - PowerPoint Processing

This skill handles any .pptx file operations including creation, reading, editing, and formatting.

## Quick Reference

| Task | Tool |
|------|------|
| Read content | `markitdown` or unpack XML |
| Edit existing | Unpack → edit XML → repack |
| Create from scratch | `pptxgenjs` |

## Reading Content

```bash
# Text extraction
pip install "markitdown[pptx]"
python -m markitdown presentation.pptx

# Visual overview
python scripts/thumbnail.py presentation.pptx
```

## Creating from Scratch

Use `pptxgenjs` library for creating presentations from scratch.

### Design Principles

**Don't create boring slides.** Plain bullets on a white background won't impress anyone.

### Color Palettes

Choose a bold, content-informed color palette:

| Theme | Primary | Secondary | Accent |
|-------|---------|-----------|--------|
| Midnight Executive | `#1E2761` | `#CADCFC` | `#FFFFFF` |
| Forest & Moss | `#2C5F2D` | `#97BC62` | `#F5F5F5` |
| Coral Energy | `#F96167` | `#F9E795` | `#2F3C7E` |
| Warm Terracotta | `#B85042` | `#E7E8D1` | `#A7BEAE` |
| Ocean Gradient | `#065A82` | `#1C7293` | `#21295C` |
| Charcoal Minimal | `#36454F` | `#F2F2F2` | `#212121` |
| Teal Trust | `#028090` | `#00A896` | `#02C39A` |

### Typography

| Header Font | Body Font |
|------------|-----------|
| Georgia | Calibri |
| Arial Black | Arial |
| Trebuchet MS | Calibri |
| Impact | Arial |
| Palatino | Garamond |

| Element | Size |
|---------|------|
| Slide title | 36-44pt |
| Section header | 20-24pt |
| Body text | 14-16pt |
| Captions | 10-12pt |

### Layout Tips

- Every slide needs a visual element
- Use asymmetric layouts
- Add large stat callouts for data
- Use icons in colored circles
- Leave breathing room (0.3-0.5" gaps)

### Avoid

- **Don't repeat the same layout** - vary across slides
- **Don't center body text** - left-align
- **Don't default to blue** - pick topic-appropriate colors
- **Don't create text-only slides** - add images/icons/charts
- **Never use accent lines under titles** - use whitespace instead

## Dependencies

```bash
pip install "markitdown[pptx]"
pip install Pillow
npm install -g pptxgenjs
```
