#!/usr/bin/env node
/**
 * Claude Code Hook 共享工具模块
 * 提供 stdin 读取、文件路径提取、二进制检测、安全文件读取
 */
const fs = require('fs');
const path = require('path');

const BINARY_EXTENSIONS = new Set([
  '.png', '.jpg', '.jpeg', '.gif', '.bmp', '.ico', '.webp',
  '.zip', '.tar', '.gz', '.rar', '.7z', '.bz2',
  '.exe', '.dll', '.so', '.dylib', '.bin',
  '.woff', '.woff2', '.ttf', '.eot', '.otf',
  '.mp3', '.mp4', '.wav', '.avi', '.mov', '.mkv', '.flv',
  '.pdf', '.doc', '.docx', '.xls', '.xlsx', '.ppt', '.pptx',
  '.class', '.o', '.obj', '.pyc', '.pyo', '.jar', '.war',
  '.sqlite', '.db', '.wasm', '.node',
]);

const MAX_FILE_SIZE = 2 * 1024 * 1024; // 2MB

/**
 * 从 stdin 读取并解析 Hook payload
 * 修复: 同时支持 file_path 和 path 字段
 * @returns {Promise<{toolName: string, filePath: string, raw: object} | null>}
 */
function readPayload() {
  return new Promise((resolve) => {
    let inputData = '';
    let settled = false;
    const stdin = process.stdin;
    stdin.setEncoding('utf8');

    const cleanup = () => {
      stdin.removeAllListeners();
    };

    const safeResolve = (value) => {
      if (!settled) {
        settled = true;
        resolve(value);
      }
    };

    // 15秒超时兜底：高负载下Claude Code发送payload可能延迟
    const timeout = setTimeout(() => {
      cleanup();
      console.error('[hook-utils] readPayload TIMEOUT: no stdin data received within 15s — hook will be skipped');
      safeResolve(null);
    }, 15000);

    stdin.on('data', (chunk) => { if (!settled) inputData += chunk; });
    stdin.on('end', () => {
      clearTimeout(timeout);
      cleanup();
      try {
        const payload = JSON.parse(inputData);
        // 修复关键 Bug: Claude Code 工具使用 file_path 而非 path
        const filePath = payload.tool_input?.file_path || payload.tool_input?.path;
        if (!filePath) {
          safeResolve(null);
          return;
        }
        safeResolve({ toolName: payload.tool_name, filePath, raw: payload });
      } catch (e) {
        console.error('Hook payload parse error:', e.message);
        safeResolve(null);
      }
    });
    stdin.on('error', (e) => {
      if (settled) return;
      clearTimeout(timeout);
      cleanup();
      console.error('Stdin read error:', e.message);
      safeResolve(null);
    });
  });
}

/**
 * 判断文件是否为二进制格式
 */
function isBinaryFile(filePath) {
  const ext = path.extname(filePath).toLowerCase();
  return BINARY_EXTENSIONS.has(ext);
}

/**
 * 安全读取文件内容（带大小限制）
 * @returns {string | null}
 */
function readFileContent(filePath) {
  try {
    if (!fs.existsSync(filePath)) return null;
    const stat = fs.statSync(filePath);
    if (stat.size > MAX_FILE_SIZE) {
      console.error(`File too large (${Math.round(stat.size / 1024)}KB), skipping: ${filePath}`);
      return null;
    }
    return fs.readFileSync(filePath, 'utf8');
  } catch (e) {
    console.error('File read error:', e.message);
    return null;
  }
}

/**
 * 获取文件扩展名（小写）
 */
function getFileExt(filePath) {
  return path.extname(filePath).toLowerCase();
}

/**
 * 从 tool_input 提取即将写入的内容（用于 PreToolUse Hook）
 * Write 工具: tool_input.content
 * Edit 工具: tool_input.new_string
 * @param {object} raw - 完整 payload
 * @returns {string | null}
 */
function extractNewContent(raw) {
  const input = raw?.tool_input;
  if (!input) return null;
  // Write 工具提交完整文件内容
  if (typeof input.content === 'string') return input.content;
  // Edit 工具提交替换片段
  if (typeof input.new_string === 'string') return input.new_string;
  return null;
}

/**
 * 追加日志到 hooks/logs/ 目录（按天切割，单文件上限 1MB 自动轮转）
 * @param {string} hookName - 日志文件前缀，如 'secrets' 或 'debug'
 * @param {string} line - 日志行内容
 */
function appendHookLog(hookName, line) {
  const logsDir = path.join(__dirname, 'logs');
  try {
    if (!fs.existsSync(logsDir)) fs.mkdirSync(logsDir, { recursive: true });
    const date = new Date().toISOString().slice(0, 10);
    const logFile = path.join(logsDir, `${hookName}-${date}.log`);
    const stat = fs.existsSync(logFile) ? fs.statSync(logFile) : null;
    // 超过 1MB 轮转
    if (stat && stat.size > 1024 * 1024) {
      const rotated = path.join(logsDir, `${hookName}-${date}-${Math.floor(Date.now() / 1000)}.log`);
      fs.renameSync(logFile, rotated);
    }
    fs.appendFileSync(logFile, line + '\n', 'utf8');
  } catch (_) {
    // 日志写入失败不影响 hook 主流程
  }
}

module.exports = { readPayload, isBinaryFile, readFileContent, getFileExt, extractNewContent, appendHookLog };
