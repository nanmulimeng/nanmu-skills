package mcp

import (
	"encoding/json"
	"log"
	"net/http"
)

type ToolHandler func(args json.RawMessage) (interface{}, error)

type Server struct {
	addr     string
	handlers map[string]ToolHandler
	toolDefs []ToolDef
	authKey  string
}

func NewServer(authKey string) *Server {
	s := &Server{
		handlers: make(map[string]ToolHandler),
		authKey:  authKey,
	}
	return s
}

func (s *Server) RegisterTool(name, desc string, props map[string]Property, required []string, handler ToolHandler) {
	s.toolDefs = append(s.toolDefs, ToolDef{
		Name:        name,
		Description: desc,
		InputSchema: InputSchema{
			Type:       "object",
			Properties: props,
			Required:   required,
		},
	})
	s.handlers[name] = handler
}

func (s *Server) ListenAndServe(addr string) error {
	s.addr = addr
	mux := http.NewServeMux()
	mux.HandleFunc("/mcp", s.handleMCP)
	mux.HandleFunc("/health", s.handleHealth)
	log.Printf("MCP server starting on %s", addr)
	return http.ListenAndServe(addr, mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       "ok",
		"tools_count":  len(s.toolDefs),
		"tools":        toolNames(s.toolDefs),
	})
}

func toolNames(defs []ToolDef) []string {
	names := make([]string, len(defs))
	for i, d := range defs {
		names[i] = d.Name
	}
	return names
}

func (s *Server) handleMCP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	key := r.Header.Get("X-Skill-Key")
	if s.authKey != "" && key != s.authKey {
		writeError(w, nil, -32001, "未授权：X-Skill-Key 无效")
		return
	}

	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, nil, -32700, "JSON解析错误")
		return
	}

	log.Printf("MCP method: %s", req.Method)

	switch req.Method {
	case "initialize":
		s.handleInitialize(w, req)
	case "tools/list":
		s.handleToolsList(w, req)
	case "tools/call":
		s.handleToolsCall(w, req)
	case "notifications/initialized":
		// No response needed for notifications
	default:
		writeError(w, req.ID, -32601, "未知方法: "+req.Method)
	}
}

func (s *Server) handleInitialize(w http.ResponseWriter, req JSONRPCRequest) {
	writeResult(w, req.ID, InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities: Capabilities{
			Tools: &ToolsCapability{ListChanged: false},
		},
		ServerInfo: ServerInfo{
			Name:    "nanmu-skill-mcp",
			Version: "1.0.0",
		},
	})
}

func (s *Server) handleToolsList(w http.ResponseWriter, req JSONRPCRequest) {
	writeResult(w, req.ID, ToolsListResult{Tools: s.toolDefs})
}

func (s *Server) handleToolsCall(w http.ResponseWriter, req JSONRPCRequest) {
	var params ToolsCallParams
	if err := json.Unmarshal(req.Params, &params); err != nil {
		writeError(w, req.ID, -32602, "参数解析错误")
		return
	}

	handler, ok := s.handlers[params.Name]
	if !ok {
		writeError(w, req.ID, -32602, "未知工具: "+params.Name)
		return
	}

	result, err := handler(params.Arguments)
	if err != nil {
		writeResult(w, req.ID, ToolsCallResult{
			Content: []ContentItem{{Type: "text", Text: err.Error()}},
			IsError: true,
		})
		return
	}

	text, _ := json.MarshalIndent(result, "", "  ")
	writeResult(w, req.ID, ToolsCallResult{
		Content: []ContentItem{{Type: "text", Text: string(text)}},
	})
}

func writeResult(w http.ResponseWriter, id json.RawMessage, result interface{}) {
	json.NewEncoder(w).Encode(JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	})
}

func writeError(w http.ResponseWriter, id json.RawMessage, code int, msg string) {
	json.NewEncoder(w).Encode(JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &JSONRPCError{Code: code, Message: msg},
	})
}
