#!/usr/bin/env node
/**
 * Claude Code Hook (PreToolUse): 写入前检测密钥和凭据
 * 检测到密钥时：
 *   - 非配置文件 → exit(2) 阻断操作
 *   - 配置文件 → stdout 警告，exit(0) 放行
 * 仅处理 Write / Edit 工具，跳过其他工具
 */
const path = require('path');
const { readPayload, isBinaryFile, extractNewContent, appendHookLog } = require(path.join(__dirname, 'hook-utils'));

// 配置文件白名单（允许存放密钥，仅警告不阻断）
const CONFIG_FILES = new Set([
  '.env', '.env.local', '.env.development', '.env.production',
  '.env.test', '.env.example', '.env.sample',
  'settings.json', 'settings.local.json', 'config.json',
]);

// 已知格式的密钥模式（高置信度，低误报）
const KNOWN_KEY_PATTERNS = [
  // OpenAI (sk-) 和 Anthropic (sk-ant-) 分开匹配，避免误报
  { pattern: /\bsk-ant-api03-[a-zA-Z0-9_-]{80,}/g, name: 'Anthropic API Key' },
  { pattern: /\bsk-[a-zA-Z0-9]{20}T3[a-zA-Z0-9_-]{16,}/g, name: 'OpenAI API Key (project)' },
  { pattern: /\bsk-[a-f0-9]{32,}/g, name: 'OpenAI API Key (legacy hex)' },
  { pattern: /\bAKIA[0-9A-Z]{16}\b/g, name: 'AWS Access Key ID' },
  { pattern: /\bASIA[0-9A-Z]{16}\b/g, name: 'AWS Temporary (STS) Key ID' },
  { pattern: /\bghp_[a-zA-Z0-9]{36}/g, name: 'GitHub Personal Token' },
  { pattern: /\bgho_[a-zA-Z0-9]{36}/g, name: 'GitHub OAuth Token' },
  { pattern: /\bghu_[a-zA-Z0-9]{36}/g, name: 'GitHub User Token' },
  { pattern: /\bghs_[a-zA-Z0-9]{36}/g, name: 'GitHub App Token' },
  { pattern: /\bglpat-[a-zA-Z0-9\-]{20,}/g, name: 'GitLab Personal Access Token' },
  { pattern: /\bglptt-[a-zA-Z0-9\-]{20,}/g, name: 'GitLab Pipeline Trigger Token' },
  { pattern: /\bxox[bpras]-[a-zA-Z0-9-]{10,}/g, name: 'Slack Token' },
  { pattern: /\bAIza[a-zA-Z0-9_-]{35}/g, name: 'Google API Key' },
  { pattern: /\bsk_live_[a-zA-Z0-9]{24,}/g, name: 'Stripe Secret Key' },
  { pattern: /\bpk_live_[a-zA-Z0-9]{24,}/g, name: 'Stripe Publishable Key' },
  { pattern: /\brk_live_[a-zA-Z0-9]{24,}/g, name: 'Stripe Restricted Key' },
  { pattern: /\bnpm_[a-zA-Z0-9]{36,}/g, name: 'npm Access Token' },
  // Fix: removed \b prefix because '-' is not a word character
  { pattern: /-----BEGIN\s+(?:RSA |DSA |EC |OPENSSH )?PRIVATE\s+KEY/g, name: 'Private Key' },
];

function getBasename(filePath) {
  return filePath.split(/[\\/]/).pop() || '';
}

function isConfigFile(filePath) {
  return CONFIG_FILES.has(getBasename(filePath));
}

function findSecrets(content) {
  const found = [];
  for (const { pattern, name } of KNOWN_KEY_PATTERNS) {
    const re = new RegExp(pattern.source, pattern.flags);
    const matches = content.match(re);
    if (matches) found.push(`- ${name} (${matches.length})`);
  }
  return found;
}

(async () => {
  const payload = await readPayload();
  if (!payload) process.exit(0);

  const { toolName, filePath, raw } = payload;

  // Fix: 仅处理 Write / Edit 工具
  if (toolName !== 'Write' && toolName !== 'Edit') {
    process.exit(0);
  }

  if (isBinaryFile(filePath)) process.exit(0);

  // 从 tool_input 提取即将写入的内容（PreToolUse 模式）
  const content = extractNewContent(raw);
  if (!content) process.exit(0);

  const found = findSecrets(content);

  if (found.length) {
    const logLine = `[${new Date().toISOString()}] ${isConfigFile(filePath) ? 'WARN' : 'BLOCK'} secrets in ${filePath}: ${found.join(', ')}`;
    appendHookLog('secrets', logLine);
    if (isConfigFile(filePath)) {
      // 配置文件：仅警告，不阻断
      console.log(`⚠️ Secrets detected in config file ${filePath} — ensure this is not committed to version control:\n${found.join('\n')}`);
      process.exit(0);
    } else {
      // 非配置文件：阻断写入，exit(2) 阻止工具执行
      console.error(`🔐 BLOCKED — Secrets detected in ${filePath}:\n${found.join('\n')}\nOperation aborted to prevent credential leakage.`);
      process.exit(2);
    }
  }

  process.exit(0);
})();
