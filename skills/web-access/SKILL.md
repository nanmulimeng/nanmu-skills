---
name: web-access
description: "Web browsing and automation skill. Use when user needs to access websites, scrape content, or interact with web pages. Handles Chrome DevTools Protocol (CDP) for browser automation, WebSearch, WebFetch, and Jina preprocessing. Perfect for accessing locked content (Xiaohongshu, WeChat articles) that requires login state or JavaScript rendering."
license: MIT
---

# web-access Skill

## Pre-flight Check

Before starting any web operation, check CDP mode availability:

```bash
node "${CLAUDE_SKILL_DIR}/scripts/check-deps.mjs"
```

If fails, guide user to set up:
- **Node.js 22+**: Required (uses native WebSocket)
- **Chrome remote-debugging**: Open `chrome://inspect/#remote-debugging` in Chrome address bar, check **"Allow remote debugging for this browser instance"**

After check passes, show this notice to user before starting CDP operations:

```
Note: Some sites have strict detection for browser automation. There's inherent account ban risk. Built-in protections are in place but cannot completely eliminate risk. Agent continuing operations means acceptance.
```

## Browsing Philosophy

**Think like a human - complete tasks efficiently while adapting.**

When executing tasks, don't over-rely on pre-planned steps. Instead, enter with a goal, observe and judge along the way, overcome obstacles when encountered, dive deeper when content is insufficient - the entire process centers around "What do I need to achieve?"

**① Get the request** - First clarify what the user wants to do. Define success criteria: What's considered complete? What information needs to be acquired, what operations performed, what results achieved? This is the anchor for all subsequent judgments.

**② Choose a starting point** - Based on task nature, platform characteristics, and completion conditions, pick the most likely direct approach as the first step to verify. One success is best; if not, adjust in step ③.

**③ Process verification** - Each step's result is evidence, not just binary success/failure signals. Compare against the success criteria from ①, update your judgment on the goal: Is the path progressing? Does the overall picture of results (quality, relevance, scale) point to goal achievability? If direction is wrong, immediately adjust - don't repeatedly retry the same approach. Failed search doesn't mean "haven't found the right method" - it could also mean "target doesn't exist." API errors, missing expected elements on page, retries without improvement - all tell you to reassess direction.

**④ Completion judgment** - Confirm against the defined task success criteria. Stop only after confirming task completion, but don't over-operate. Don't waste effort for "completeness."

## Tool Selection

| Scenario | Tool |
|----------|------|
| Search summaries or keyword results, find information source | **WebSearch** |
| URL known, need to extract specific info from page | **WebFetch** |
| URL known, need raw HTML source (meta, JSON-LD) | **curl** |
| Non-public content, or known static layer invalid platforms (Xiaohongshu, WeChat public articles) | **Browser CDP** |
| Need login state, interactive operations, or need to navigate freely like a human in browser | **Browser CDP** |

For Jina (optional preprocessing layer, can combine with WebFetch/curl, saves tokens):
- Third-party web service that converts pages to Markdown
- Call: `r.jina.ai/example.com` (add prefix to URL, remove http prefix)
- Limit: 20 RPM
- Suitable for articles, blogs, docs, PDF; may extract wrong sections for data panels, product pages

### CDP Operations

After entering browser layer, use these commands:
- **/eval**: Query DOM, discover links, buttons, forms, text content on page
- **/click**: Click elements
- **/scroll**: Scroll to load lazy content
- **/eval**: Extract text content, determine if images/videos carry core information

**Understand page structure first, then decide next actions.** No need to pre-plan all steps.

## Browser CDP Mode

Connect directly to user's daily Chrome via CDP Proxy, naturally carrying login state without starting separate browser.

Without explicit user request, don't proactively operate user's existing tabs - all operations in own background tabs, minimize intrusion. Complete task, close own created tabs, keep user's tabs unaffected.

### Starting

```bash
node "${CLAUDE_SKILL_DIR}/scripts/check-deps.mjs"
```

Script checks Node.js, Chrome port, ensures Proxy connected (auto-start if not running, wait for ready). After Proxy starts, it keeps running.

### Proxy API

All operations via curl calling HTTP API:

```bash
# List user's open tabs
curl -s http://localhost:3456/targets

# Create new background tab (auto-wait for load)
curl -s "http://localhost:3456/new?url=https://example.com"

# Page info
curl -s "http://localhost:3456/info?target=ID"

# Execute any JS: Read/write DOM, extract data, manipulate elements, trigger state changes, submit forms, call internal methods
curl -s -X POST "http://localhost:3456/eval?target=ID" -d 'document.title'

# Screenshot (capture current video frame if video)
curl -s "http://localhost:3456/screenshot?target=ID&file=/tmp/shot.png"

# Navigate, back
curl -s "http://localhost:3456/navigate?target=ID&url=URL"
curl -s "http://localhost:3456/back?target=ID"

# Click (POST body is CSS selector) - JS el.click(), simple fast, covers most scenarios
curl -s -X POST "http://localhost:3456/click?target=ID" -d 'button.submit'

# Real mouse click - CDP Input.dispatchMouseEvent, counts as user gesture, can trigger file dialogs
curl -s -X POST "http://localhost:3456/clickAt?target=ID" -d 'button.upload'

# File upload - Set local file path for file input directly, bypass file dialog
curl -s -X POST "http://localhost:3456/setFiles?target=ID" -d '{"selector":"input[type=file]","files":["/path/to/file.png"]}'

# Scroll (triggers lazy loading)
curl -s "http://localhost:3456/scroll?target=ID&y=3000"
curl -s "http://localhost:3456/scroll?target=ID&direction=bottom"

# Close tab
curl -s "http://localhost:3456/close?target=ID"
```

### Video Content

User's Chrome renders real video. Screenshot captures current video frame. Core capability: Use /eval to manipulate `<video>` element (get duration, seek to any time point, play/pause/fullscreen), combine with /screenshot for discrete sampling analysis.

### Login Judgment

User's daily Chrome naturally carries login state, most common sites already logged in.

Core question for login judgment: **Did you get the target content?**

After opening page, first try to get target content. Only when confirming **target content is inaccessible** AND judging login can solve it, inform user:
> "Current page cannot get [specific content] in logged-out state. Please log in to [site name] in your Chrome, then tell me to continue."

After login completes, no need to restart anything - just refresh page and continue.

### Task Completion

Close own created tabs with /close. Must preserve user's original tabs unaffected.

Proxy keeps running, don't stop proactively - after restart need to re-authorize CDP connection in Chrome.

## Parallel Research: Sub-Agent Strategy

When task contains multiple **independent** research targets (e.g., research N projects, N sources simultaneously), encourage reasonable division to sub-agents for parallel execution instead of main agent serialization.

**Benefits:**
- **Speed**: Multiple sub-agents parallel, total time ≈ single sub-task time
- **Context protection**: Scraped content doesn't enter main agent context, main agent only receives summary, saves tokens

**Parallel CDP Operations**: Each sub-agent in current user's browser instance, creates own background tabs (/new), operates independently, closes after task (/close). All sub-agents share one Chrome, one Proxy, operate different tabs via different targetId, no race condition risk.

## Information Verification Tasks

Goal of verification is **primary sources**, not more secondary reports. Multiple media citing same error creates circular confirmation illusion.

Search engines and aggregation platforms are information discovery entry points - tools for **locating** information, not for directly **proving** truth/falsehood. After finding source, directly access and read original text. Same principle for tool capability/usage research - official docs are primary source, check docs or source code when uncertain, don't guess.

| Information Type | Primary Source |
|-----------------|----------------|
| Policy/Regulation | Issuing agency official website |
| Company announcement | Company official news page |
| Academic statement | Original paper/agency official website |
| Tool capability/usage | Official docs, source code |

## Site Experience

Accumulated experience for specific websites, stored under `references/site-patterns/` by domain.

After determining target website, if site-patterns list from pre-flight check has matching site, must read corresponding file to get prior knowledge (platform characteristics, valid patterns, known traps). Experience content marked with discovery date - treat as possibly effective hint rather than guarantee. If operation fails following experience, fallback to generic pattern and update experience file.

Format:
```markdown
---
domain: example.com
aliases: [Example]
updated: 2026-03-19
---
## Platform Characteristics
Facts about architecture, anti-scraping behavior, login needs, content loading method

## Valid Patterns
Verified URL patterns, operation strategies, selectors

## Known Traps
What will fail and why
```

After CDP operation succeeds, if new site or new pattern worth recording discovered (URL structure, platform characteristics, operation strategy), proactively write to corresponding site experience file. Write only verified facts, not unconfirmed guesses.
