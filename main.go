package main

import (
	"errors"
	"fmt"
	"log"
	"monica-proxy/internal/apiserver"
	"monica-proxy/internal/config"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()
	if cfg.MonicaCookie == "" {
		log.Fatal("MONICA_COOKIE environment variable is required")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// 注册路由
	apiserver.RegisterRoutes(e)
	
	// 使用配置的端口启动服务
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	log.Printf("Starting server on %s", addr)
	if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("start server error: %v", err)
	}
}
