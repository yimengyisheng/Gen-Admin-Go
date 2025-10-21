
package main

import (
	"ai_admin_project/config"
	"ai_admin_project/internal/handler"
	"ai_admin_project/internal/middleware"
	"ai_admin_project/internal/model"
	"ai_admin_project/internal/repository"
	"ai_admin_project/internal/service"
	"ai_admin_project/pkg/logger"
	"ai_admin_project/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Init config
	config.Init()

	// Init logger
	logger.Init()

	// Init translator
	if err := utils.InitTranslator(); err != nil {
		log.Fatalf("Failed to init translator: %v", err)
	}

	// Init database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Product{})

	// Init repository
	userRepo := repository.NewUserRepository(db)

	// Init service
	userService := service.NewUserService(userRepo)

	// Init handler
	userHandler := handler.NewUserHandler(userService)

	// Init Gin engine
	r := gin.Default()

	// Public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)


	// Init Product service
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Protected routes
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", userHandler.GetProfile)

		productGroup := authGroup.Group("/products")
		{
			productGroup.POST("/create", productHandler.Create)
			productGroup.GET("/:id", productHandler.Get)
			productGroup.POST("/update", productHandler.Update)
			productGroup.POST("/delete", productHandler.Delete)
			productGroup.GET("/", productHandler.List)
		}
	}

	// Start server
	port := fmt.Sprintf(":%d", config.Conf.Server.Port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
