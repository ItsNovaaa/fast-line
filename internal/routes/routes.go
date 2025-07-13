package routes

import (
    "database/sql"
    "fast-line/internal/config"
    "fast-line/internal/handlers"
    "fast-line/internal/middleware"
    "fast-line/internal/repository"
    "fast-line/internal/services"
    
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {
    // Add middleware
    router.Use(middleware.CORS())
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    
    // Initialize repositories
    userRepo := repository.NewUserRepository(db)
    circuitRepo := repository.NewCircuitRepository(db)
    
    // Initialize services
    authService := services.NewAuthService(userRepo, cfg.JWTSecret)
    circuitService := services.NewCircuitService(circuitRepo)
    
    // Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)

    //Circuit Handlers
    circuitHandler := handlers.NewCircuitHandler(circuitService)
    
    // API routes group
    api := router.Group("/api/v1")
    {
        // Health check
        api.GET("/health", func(c *gin.Context) {
            c.JSON(200, gin.H{"status": "hai"})
        })
        
        // Auth routes
        auth := api.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
        }

        circuit := api.Group("/circuits")
        circuit.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        {
            circuit.GET("/test", func(c *gin.Context) {
                c.JSON(200, gin.H{"status": "OK"})
            })
            circuit.POST("/create", circuitHandler.Create)
            circuit.PUT("/update/:id", circuitHandler.Update)
            circuit.GET("/get/:id", circuitHandler.Get)
            circuit.GET("/list", circuitHandler.List)
            circuit.PUT("/delete/:id", circuitHandler.Delete)
        }
        
        // Protected routes (add middleware later)
        // protected := api.Group("/")
        // protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        // {
        //     protected.GET("/profile", authHandler.GetProfile)
        // }
    }
}
