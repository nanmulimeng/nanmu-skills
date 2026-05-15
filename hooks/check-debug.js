#!/usr/bin/env node
/**
 * Claude Code Hook: 检测文件中的调试代码 (PreToolUse)
 * 从 stdin payload 读取即将写入的新内容，支持语言感知匹配
 */
const path = require('path');
const { readPayload, extractNewContent, getFileExt, appendHookLog } = require(path.join(__dirname, 'hook-utils'));

// 明确的调试断点/语句（仅保留不会在正式代码中出现的模式）
const DEBUG_PATTERNS = [
  // JavaScript / TypeScript
  { pattern: /\bdebugger\s*;?/g, name: 'debugger', exts: ['.js', '.jsx', '.ts', '.tsx', '.mjs', '.cjs'] },
  // Python
  { pattern: /\bbreakpoint\s*\(/g, name: 'breakpoint()', exts: ['.py', '.pyw'] },
  { pattern: /\bpdb\.set_trace\s*\(/g, name: 'pdb.set_trace()', exts: ['.py', '.pyw'] },
  { pattern: /\bipdb\s*\(/g, name: 'ipdb', exts: ['.py', '.pyw'] },
  // Rust
  { pattern: /\bdbg!\s*\(/g, name: 'dbg!()', exts: ['.rs'] },
  // PHP
  { pattern: /\bvar_dump\s*\(/g, name: 'var_dump()', exts: ['.php'] },
  { pattern: /\bdd\s*\(/g, name: 'dd()', exts: ['.php'] },
  { pattern: /\bdump\s*\(/g, name: 'dump()', exts: ['.php'] },
  // Dart / Flutter
  { pattern: /\bdebugPrint\s*\(/g, name: 'debugPrint()', exts: ['.dart'] },
  // Go — fmt/log debug prints
  { pattern: /\b(log|fmt)\.(Printf|Println|Print)\s*\([^)]*"[Dd][Ee][Bb][Uu][Gg]/g, name: 'log/fmt.Printf("DEBUG...")', exts: ['.go'] },
  // Java — System.out/err debug
  { pattern: /\bSystem\.(out|err)\.(println|printf|print)\s*\([^)]*"[Dd][Ee][Bb][Uu][Gg]/g, name: 'System.out.println("DEBUG...")', exts: ['.java'] },
  // C/C++ — printf/cout debug
  { pattern: /\bprintf\s*\([^)]*"[Dd][Ee][Bb][Uu][Gg]/g, name: 'printf("DEBUG...")', exts: ['.c', '.cpp', '.cc', '.cxx', '.h', '.hpp'] },
];

(async () => {
  const payload = await readPayload();
  if (!payload) process.exit(0);

  const { toolName, filePath, raw } = payload;

  // 仅处理 Write 和 Edit 工具
  if (toolName !== 'Write' && toolName !== 'Edit') {
    process.exit(0);
  }

  const content = extractNewContent(raw);
  if (!content) process.exit(0);

  const ext = getFileExt(filePath);
  const found = [];

  for (const { pattern, name, exts } of DEBUG_PATTERNS) {
    // exts=null 表示对所有语言生效
    if (exts !== null && !exts.includes(ext)) continue;
    // 克隆正则以避免 lastIndex 残留
    const re = new RegExp(pattern.source, pattern.flags);
    const matches = content.match(re);
    if (matches) found.push(`- ${name} (${matches.length})`);
  }

  if (found.length) {
    const logLine = `[${new Date().toISOString()}] WARN debug in ${filePath}: ${found.join(', ')}`;
    appendHookLog('debug', logLine);
    console.log(`⚠️ Debug code detected in ${filePath} (tool: ${toolName}):\n${found.join('\n')}`);
  }

  process.exit(0);
})();
