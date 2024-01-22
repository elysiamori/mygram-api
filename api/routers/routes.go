package routers

import (
	"github.com/elysiamori/mygram-api/api/handler"
	"github.com/elysiamori/mygram-api/api/middlewares"
	"github.com/elysiamori/mygram-api/api/repositories/mygram"
	"github.com/elysiamori/mygram-api/api/services"
	"github.com/gofiber/fiber/v2"
)

func MyGramRoutes(app *fiber.App, userService *services.UserServiceImpl, photoService *services.PhotoServiceImpl,
	commentService *services.CommentServiceImpl, socialMediaService *services.SocialMediaServiceImpl) {

	userRepository := mygram.UserRepositoryImpl{}
	userHandler := handler.NewUserHandler(*userService, userRepository)

	photoHandler := handler.NewPhotoHandler(*photoService, *userService)

	commentHandler := handler.NewCommentHandler(*commentService)

	socialMediaHandler := handler.NewSocialMediaHandler(*socialMediaService)

	// users
	userGroup := app.Group("/users")
	{
		userGroup.Post("/register", userHandler.RegisterUser)
		userGroup.Post("/login", userHandler.LoginUser)
		userGroup.Post("/loginusername", userHandler.LoginUserWithUsername)
		userGroup.Put("/", middlewares.JwtAuthMiddleware(), userHandler.UpdateUser)
		userGroup.Delete("/", middlewares.JwtAuthMiddleware(), userHandler.DeleteUser)
	}

	profileGroup := app.Group("/profile")
	{
		profileGroup.Use(middlewares.JwtAuthMiddleware())
		profileGroup.Get("/", userHandler.ProfileUser)

	}

	// photos
	photoGroup := app.Group("/photos")
	{
		photoGroup.Use(middlewares.JwtAuthMiddleware())
		photoGroup.Post("/", photoHandler.UploadPhoto)
		photoGroup.Get("/", photoHandler.GetAllPhoto)
		photoGroup.Get("/:id", photoHandler.GetPhotoByID)
		photoGroup.Put("/:id", photoHandler.UpdatePhoto)
		photoGroup.Delete("/:id", photoHandler.DeletePhoto) // masih problem
	}

	// get api
	apiGroup := app.Group("/api")
	{
		apiGroup.Get("/", photoHandler.GetAPI)
	}

	// comments
	commentGroup := app.Group("/comments")
	{
		commentGroup.Use(middlewares.JwtAuthMiddleware())
		commentGroup.Post("/", commentHandler.PostComment)
		commentGroup.Get("/", commentHandler.GetAllComment)
		commentGroup.Get("/:id", commentHandler.GetCommentByID)
		commentGroup.Put("/:id", commentHandler.UpdateComment)
		commentGroup.Delete("/:id", commentHandler.DeleteComment)
	}

	// social media
	socialMediaGroup := app.Group("/socialmedias")
	{
		socialMediaGroup.Use(middlewares.JwtAuthMiddleware())
		socialMediaGroup.Post("/", socialMediaHandler.PostSocialMedia)
		socialMediaGroup.Get("/", socialMediaHandler.GetAllSocialMedia)
	}
}
