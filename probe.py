import paramiko
c = paramiko.SSHClient()
c.set_missing_host_key_policy(paramiko.AutoAddPolicy())
c.connect('123.56.223.97', username='nanmu', password='nanmu1679969808', timeout=10)
PATH = 'export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin; '
cmds = [
    'ps aux | grep nanmu',
    'systemctl list-units | grep nanmu',
    'ss -tlnp | grep 3456',
    'curl -s http://127.0.0.1:3456/mcp 2>/dev/null | head -5 || echo no-response',
    'which go && go version',
    'ls /home/nanmu/nanmu-skill-mcp/bin/ 2>/dev/null',
    'cat /home/nanmu/go.mod 2>/dev/null || echo no-go-mod',
]
for cmd in cmds:
    print(f'=== {cmd} ===')
    _, o, e = c.exec_command(PATH + cmd, timeout=10)
    out = o.read().decode()[:500]
    err = e.read().decode().strip()
    if out.strip(): print(out)
    if err: print(f'  ERR: {err}')
c.close()
