package main

import (
    "log"
    "fast-line/internal/config"
    "fast-line/internal/database"
    "fast-line/internal/routes"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

   cfg := config.Load()
    
    // Initialize database
    db, err := database.NewConnection(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()
    
    // Initialize Gin router
    router := gin.Default()
    
    // Setup routes
    routes.SetupRoutes(router, db, cfg)
    
    // Start server
    log.Printf("Server starting on port %s", cfg.Port)
    if err := router.Run(":" + cfg.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}