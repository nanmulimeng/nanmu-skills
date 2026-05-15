#!/usr/bin/env python3
"""Deploy skills to remote MCP server via SSH (paramiko)."""
import os
import sys
import getpass
import tarfile
import tempfile
import paramiko

HOST = "123.56.223.97"
PORT = 22
USER = "nanmu"
PASSWORD = os.environ.get("DEPLOY_PASSWORD", "")
REMOTE_DIR = "/home/nanmu/nanmu-skill-mcp"
SERVICE = "nanmu-skill-mcp"
LOCAL_SKILLS = os.path.join(os.path.dirname(os.path.abspath(__file__)), "skills")
PATH_PREFIX = "export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin; "

def get_ssh_client():
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    key_path = os.path.expanduser("~/.ssh/id_rsa")
    try:
        if os.path.exists(key_path):
            key = paramiko.RSAKey.from_private_key_file(key_path)
            client.connect(HOST, port=PORT, username=USER, pkey=key, timeout=10)
            return client
    except paramiko.ssh_exception.AuthenticationException:
        pass
    pwd = PASSWORD or getpass.getpass(f"Password for {USER}@{HOST}: ")
    client.connect(HOST, port=PORT, username=USER, password=pwd, timeout=10)
    return client

def run_remote(client, cmd, timeout=30):
    print(f"  $ {cmd}")
    _, stdout, stderr = client.exec_command(PATH_PREFIX + cmd, timeout=timeout)
    out = stdout.read().decode()
    err = stderr.read().decode()
    code = stdout.channel.recv_exit_status()
    if out.strip():
        print(f"  {out.strip()}")
    if err.strip():
        print(f"  [stderr] {err.strip()}")
    return code, out, err

def main():
    if not os.path.isdir(LOCAL_SKILLS):
        print(f"skills/ not found at {LOCAL_SKILLS}")
        sys.exit(1)

    # Pack
    tmp = tempfile.NamedTemporaryFile(suffix=".tar.gz", delete=False)
    tmp.close()
    print(f"Packing {LOCAL_SKILLS} -> {tmp.name}")
    with tarfile.open(tmp.name, "w:gz") as tar:
        tar.add(LOCAL_SKILLS, arcname="skills")
    size_kb = os.path.getsize(tmp.name) // 1024
    print(f"Archive: {size_kb}KB")

    # Connect
    print(f"Connecting {USER}@{HOST}:{PORT}...")
    client = get_ssh_client()
    print("Connected.")

    # Upload
    remote_tar = f"/tmp/skills-deploy-{os.getpid()}.tar.gz"
    print(f"Uploading -> {remote_tar}")
    sftp = client.open_sftp()
    sftp.put(tmp.name, remote_tar)
    sftp.close()
    print("Uploaded.")

    # Deploy: extract + restart systemd service
    print("Deploying...")
    run_remote(client, f"cd {REMOTE_DIR} && tar xzf {remote_tar} && rm -f {remote_tar}")
    run_remote(client, f"sudo systemctl restart {SERVICE}")

    import time
    time.sleep(2)

    code, out, _ = run_remote(client, f"sudo systemctl is-active {SERVICE}")
    status = out.strip()
    print(f"Service {SERVICE}: {status}")

    # Verify port listening
    run_remote(client, "ss -tlnp | grep 3456")

    client.close()
    os.unlink(tmp.name)

    if status == "active":
        print("Deploy succeeded.")
    else:
        print("WARNING: service not active after deploy.")
        sys.exit(1)

if __name__ == "__main__":
    main()
