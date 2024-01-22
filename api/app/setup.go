package app

import (
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/api/routers"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := mygram.NewUserRepository(db)
	userService := services.NewUserService(*userRepository)

	photoRepository := mygram.NewPhotoRepository(db)
	photoService := services.NewPhotoService(*photoRepository)

	commentRepository := mygram.NewCommentRepository(db)
	commentService := services.NewCommentService(*commentRepository)

	socialMediaRepository := mygram.NewSMRepository(db)
	socialMediaService := services.NewSocialMediaService(*socialMediaRepository)

	// untuk mengizinkan CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// untuk menampilkan log
	app.Use(logger.New())

	routers.MyGramRoutes(app, userService, photoService, commentService, socialMediaService)
}
