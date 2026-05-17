import paramiko
c = paramiko.SSHClient()
c.set_missing_host_key_policy(paramiko.AutoAddPolicy())
c.connect('123.56.223.97', username='nanmu', password='nanmu1679969808', timeout=10)
PATH = 'export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin; '
# Write a test program that runs the actual route logic
cmd = r"""cat > /tmp/route-test.go << 'GOEOF'
package main

import (
	"encoding/json"
	"fmt"
	"nanmu-skill-mcp/registry"
	"nanmu-skill-mcp/router"
)

func main() {
	m, _ := registry.LoadFromDir("/home/nanmu/nanmu-skill-mcp/skills")
	fmt.Printf("Skills: %d\n", m.Count())

	queries := []string{"审计代码安全", "Go语言开发", "AI Agent RAG"}
	for _, q := range queries {
		r := router.Route(m, q)
		b, _ := json.Marshal(map[string]interface{}{
			"query": q,
			"must_count": len(r.Must),
		})
		fmt.Println(string(b))
	}

	// Debug scoring
	tokens := router.Tokenize("审计代码安全")
	fmt.Printf("Tokens: %v\n", tokens)
	sk := m.Index["coff0xc-secure-code-appsec"]
	score := router.Score(sk, tokens, map[string]bool{})
	fmt.Printf("Score for secure-code-appsec: %.1f\n", score)
}
GOEOF
cd /tmp/mcp-build && /usr/local/go/bin/go run /tmp/route-test.go 2>&1"""
_, o, e = c.exec_command(PATH + cmd, timeout=30)
print('Result:', o.read().decode('utf-8', errors='replace')[:800])
print('Err:', e.read().decode('utf-8', errors='replace')[:300])
c.close()
