package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carregando configurações
	config := LoadConfig()

	// Configurando o Gin em modo release para produção
	gin.SetMode(gin.ReleaseMode)

	// Criando o router Gin
	r := gin.Default()

	// Configurando templates HTML
	r.LoadHTMLGlob("templates/*.html")

	// Configurando arquivos estáticos (CSS, JS, imagens)
	r.Static("/static", "./static")

	// Criando handlers
	handlers := NewHandlers(config)

	// Configurando rotas
	setupRoutes(r, handlers)

	// Configurando a porta
	port := ":" + config.Port
	log.Printf("🚀 Servidor FreshPoint iniciando na porta %s", port)
	log.Printf("🌐 Acesse: http://localhost%s", port)
	log.Printf("🍃 FreshPoint - Distribuição de Açaí Premium")

	// Iniciando o servidor
	if err := r.Run(port); err != nil {
		log.Fatal("❌ Erro ao iniciar o servidor:", err)
	}
}

// setupRoutes configura todas as rotas da aplicação
func setupRoutes(r *gin.Engine, handlers *Handlers) {
	// Rota principal
	r.GET("/", handlers.HomeHandler)

	// Rotas da API
	api := r.Group("/api")
	{
		api.GET("/products", handlers.ProductsAPIHandler)
		api.GET("/health", handlers.HealthCheckHandler)
	}

	// Rota para contato via WhatsApp
	r.GET("/whatsapp/:product", handlers.WhatsAppHandler)
}
