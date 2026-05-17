import paramiko, os, sys

c = paramiko.SSHClient()
c.set_missing_host_key_policy(paramiko.AutoAddPolicy())
c.connect('123.56.223.97', username='nanmu', password='nanmu1679969808', timeout=10)
PATH = 'export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin; '

# Upload new binary
print("Uploading new binary...")
sftp = c.open_sftp()
local_bin = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'nanmu-skill-mcp-linux-new')
if not os.path.exists(local_bin):
    local_bin = '/tmp/nanmu-skill-mcp-linux-new'
print(f"  Local: {local_bin} ({os.path.getsize(local_bin)//1024}KB)")
sftp.put(local_bin, '/tmp/nanmu-skill-mcp-new')
sftp.close()
print("Uploaded.")

# Deploy: backup old, replace, restart
cmds = [
    'cp /home/nanmu/nanmu-skill-mcp/bin/nanmu-skill-mcp /home/nanmu/nanmu-skill-mcp/bin/nanmu-skill-mcp.bak',
    'mv /tmp/nanmu-skill-mcp-new /home/nanmu/nanmu-skill-mcp/bin/nanmu-skill-mcp',
    'chmod +x /home/nanmu/nanmu-skill-mcp/bin/nanmu-skill-mcp',
    'sudo systemctl restart nanmu-skill-mcp',
]
for cmd in cmds:
    print(f"  $ {cmd}")
    _, o, e = c.exec_command(PATH + cmd, timeout=30)
    code = o.channel.recv_exit_status()
    err = e.read().decode().strip()
    if err: print(f"    ERR: {err}")
    if code != 0: print(f"    exit: {code}")

import time; time.sleep(2)

_, o, e = c.exec_command(PATH + 'sudo systemctl is-active nanmu-skill-mcp', timeout=10)
status = o.read().decode().strip()
print(f"Service: {status}")
c.close()
