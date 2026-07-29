package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"clients-service/internal/config"
	"clients-service/internal/handler"
	"clients-service/internal/middleware"
	"clients-service/internal/repository"
	"clients-service/internal/service"

	"github.com/gorilla/mux"
)

const apiDocumentation = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Clients Service API Documentation</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 30px;
            border-radius: 10px;
            margin-bottom: 30px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }
        .header h1 {
            margin: 0 0 10px 0;
            font-size: 2.5em;
        }
        .header p {
            margin: 0;
            opacity: 0.9;
        }
        .endpoint {
            background: white;
            border-radius: 8px;
            padding: 20px;
            margin-bottom: 20px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
            border-left: 4px solid #667eea;
        }
        .method {
            display: inline-block;
            padding: 6px 12px;
            border-radius: 4px;
            font-weight: bold;
            font-size: 0.9em;
            margin-right: 10px;
        }
        .get { background-color: #61affe; color: white; }
        .post { background-color: #49cc90; color: white; }
        .put { background-color: #fca130; color: white; }
        .delete { background-color: #f93e3e; color: white; }
        .path {
            font-family: 'Courier New', monospace;
            background-color: #f8f9fa;
            padding: 6px 12px;
            border-radius: 4px;
            font-size: 0.9em;
        }
        .description {
            margin: 15px 0;
            color: #555;
            line-height: 1.6;
        }
        .example {
            background-color: #f8f9fa;
            border: 1px solid #dee2e6;
            border-radius: 4px;
            padding: 15px;
            margin-top: 15px;
        }
        .example-title {
            font-weight: bold;
            margin-bottom: 10px;
            color: #495057;
        }
        pre {
            background-color: #282c34;
            color: #abb2bf;
            padding: 15px;
            border-radius: 4px;
            overflow-x: auto;
            margin: 0;
        }
        code {
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
        }
        .response {
            margin-top: 15px;
        }
        .response-title {
            font-weight: bold;
            color: #28a745;
            margin-bottom: 10px;
        }
        .section {
            margin-top: 40px;
        }
        .section-title {
            font-size: 1.8em;
            color: #333;
            margin-bottom: 20px;
            padding-bottom: 10px;
            border-bottom: 2px solid #667eea;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>Clients Service API</h1>
        <p>API для управления клиентами в системе Автовышкин</p>
        <p>Версия: 1.0 | Базовый URL: <code>http://localhost:8080/api</code></p>
    </div>

    <div class="section">
        <h2 class="section-title">Клиенты</h2>

        <div class="endpoint">
            <span class="method get">GET</span>
            <span class="path">/clients</span>
            <div class="description">
                Получить список всех клиентов с пагинацией
            </div>
            <div class="example">
                <div class="example-title">Параметры запроса:</div>
                <pre><code>page: номер страницы (по умолчанию 1)
perPage: количество элементов на странице (по умолчанию 10)</code></pre>
            </div>
            <div class="response">
                <div class="response-title">Пример ответа:</div>
                <pre><code>{
  "success": true,
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Иван Иванов",
      "type": "individual",
      "status": "active",
      "phone": "+79001234567",
      "email": "ivan@example.com",
      "createdAt": "2024-01-15T10:30:00Z",
      "updatedAt": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "perPage": 10,
    "total": 1,
    "totalPages": 1
  }
}</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method get">GET</span>
            <span class="path">/clients/{id}</span>
            <div class="description">
                Получить информацию о клиенте по ID
            </div>
            <div class="example">
                <div class="example-title">Пример запроса:</div>
                <pre><code>GET /api/clients/123e4567-e89b-12d3-a456-426614174000</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method get">GET</span>
            <span class="path">/clients/type/{type}</span>
            <div class="description">
                Получить клиентов по типу (individual/legal)
            </div>
            <div class="example">
                <div class="example-title">Пример запроса:</div>
                <pre><code>GET /api/clients/type/individual?page=1&perPage=10</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method get">GET</span>
            <span class="path">/clients/search</span>
            <div class="description">
                Поиск клиентов по имени, телефону или email
            </div>
            <div class="example">
                <div class="example-title">Параметры запроса:</div>
                <pre><code>query: строка поиска
page: номер страницы (по умолчанию 1)
perPage: количество элементов на странице (по умолчанию 10)</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method post">POST</span>
            <span class="path">/clients</span>
            <div class="description">
                Создать нового клиента
            </div>
            <div class="example">
                <div class="example-title">Пример запроса:</div>
                <pre><code>{
  "name": "Иван Иванов",
  "type": "individual",
  "status": "active",
  "phone": "+79001234567",
  "email": "ivan@example.com",
  "address": "г. Москва, ул. Примерная, д. 1"
}</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method put">PUT</span>
            <span class="path">/clients/{id}</span>
            <div class="description">
                Обновить информацию о клиенте
            </div>
            <div class="example">
                <div class="example-title">Пример запроса:</div>
                <pre><code>{
  "name": "Иван Иванович",
  "phone": "+79001234568",
  "status": "inactive"
}</code></pre>
            </div>
        </div>

        <div class="endpoint">
            <span class="method delete">DELETE</span>
            <span class="path">/clients/{id}</span>
            <div class="description">
                Удалить клиента
            </div>
            <div class="example">
                <div class="example-title">Пример запроса:</div>
                <pre><code>DELETE /api/clients/123e4567-e89b-12d3-a456-426614174000</code></pre>
            </div>
        </div>
    </div>

    <div class="section">
        <h2 class="section-title">Системные эндпоинты</h2>

        <div class="endpoint">
            <span class="method get">GET</span>
            <span class="path">/health</span>
            <div class="description">
                Проверка здоровья сервиса
            </div>
            <div class="response">
                <div class="response-title">Пример ответа:</div>
                <pre><code>OK</code></pre>
            </div>
        </div>
    </div>

    <div class="section">
        <h2 class="section-title">Типы клиентов</h2>
        <div class="endpoint">
            <div class="description">
                <strong>individual</strong> - Физическое лицо<br>
                <strong>legal</strong> - Юридическое лицо
            </div>
        </div>
    </div>

    <div class="section">
        <h2 class="section-title">Статусы клиентов</h2>
        <div class="endpoint">
            <div class="description">
                <strong>active</strong> - Активный<br>
                <strong>inactive</strong> - Неактивный<br>
                <strong>blocked</strong> - Заблокированный
            </div>
        </div>
    </div>

    <div class="section">
        <h2 class="section-title">Формат ошибок</h2>
        <div class="endpoint">
            <div class="response">
                <div class="response-title">Пример ответа с ошибкой:</div>
                <pre><code>{
  "success": false,
  "error": "Client not found",
  "message": "Клиент с указанным ID не найден"
}</code></pre>
            </div>
        </div>
    </div>

    <div class="section">
        <h2 class="section-title">Примечания</h2>
        <div class="endpoint">
            <div class="description">
                <ul>
                    <li>Все даты и времени возвращаются в формате ISO 8601 (UTC)</li>
                    <li>Все ID возвращаются в формате UUID</li>
                    <li>API поддерживает CORS для фронтенд-приложения</li>
                    <li>Все запросы логируются для отладки</li>
                </ul>
            </div>
        </div>
    </div>
</body>
</html>
`

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting Clients Service...")
	log.Printf("Environment: %s", cfg.Env)
	log.Printf("Database: %s:%s/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// Initialize repository
	repo, err := repository.NewClientRepository(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	// Initialize service
	clientService := service.NewClientService(repo)

	// Initialize handlers
	clientHandler := handler.NewClientHandler(clientService)

	// Create router
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware(cfg.CORS.AllowedOrigins))

	// Global OPTIONS handler for CORS preflight requests
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Client routes
	api.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST", "OPTIONS")
	api.HandleFunc("/clients", clientHandler.GetClients).Methods("GET", "OPTIONS")
	api.HandleFunc("/clients/search", clientHandler.SearchClients).Methods("GET", "OPTIONS")
	api.HandleFunc("/clients/type/{type}", clientHandler.GetClientsByType).Methods("GET", "OPTIONS")
	api.HandleFunc("/clients/{id}", clientHandler.GetClient).Methods("GET", "OPTIONS")
	api.HandleFunc("/clients/{id}", clientHandler.UpdateClient).Methods("PUT", "OPTIONS")
	api.HandleFunc("/clients/{id}", clientHandler.DeleteClient).Methods("DELETE", "OPTIONS")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// API documentation
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(apiDocumentation))
	}).Methods("GET")

	// Create server
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
