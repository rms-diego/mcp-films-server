package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rms-diego/mcp-films-server/internal/config"
	"github.com/rms-diego/mcp-films-server/internal/routes"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:    "A films server for Model Context Protocol",
		Version: "v1",
	}, nil)

	s := gin.Default()
	routes.Init(s, mcpServer)

	if err := http.ListenAndServe(":"+config.Cfg.PORT, s); err != nil {
		panic(err)
	}
}
