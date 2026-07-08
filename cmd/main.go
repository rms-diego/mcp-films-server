package main

import (
	"fmt"
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

	ms := mcp.NewServer(&mcp.Implementation{
		Name:    "A films server for Model Context Protocol",
		Version: "v1",
	}, nil)

	s := gin.Default()
	routes.Init(s, ms)

	addr := fmt.Sprintf(":%v", config.Cfg.PORT)
	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}
