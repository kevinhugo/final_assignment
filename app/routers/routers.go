package routers

import (
	_ "assignment/docs"
	"log"

	"assignment/app/handler"
	"assignment/app/middleware"

	"github.com/gin-gonic/gin"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	UserHandler := handler.NewUserHandler()
	PhotoHandler := handler.NewPhotoHandler()
	CommentHandler := handler.NewCommentHandler()
	SocialMediaHandler := handler.NewSocialMediaHandler()
	r := gin.Default()

	r.POST("/User/Register", UserHandler.RegisterUser)
	r.POST("/User/Login", UserHandler.LoginUser)
	routers := r.Group("")
	routers.Use(middleware.JwtAuthMiddleware())
	routers.DELETE("/User", UserHandler.DeleteUser)
	routers.PUT("/User", UserHandler.UpdateUser)
	routers.POST("/Photos", PhotoHandler.AddPhoto)
	routers.GET("/Photos", PhotoHandler.GetPhoto)
	routers.PUT("/Photos/:photoId", PhotoHandler.UpdatePhoto)
	routers.DELETE("/Photos/:photoId", PhotoHandler.DeletePhoto)

	routers.POST("/Comments", CommentHandler.AddComment)
	routers.GET("/Comments", CommentHandler.GetComment)
	routers.PUT("/Comments/:commentId", CommentHandler.UpdateComment)
	routers.DELETE("/Comments/:commentId", CommentHandler.DeleteComment)

	routers.POST("/SocialMedias", SocialMediaHandler.AddSocialMedia)
	routers.GET("/SocialMedias", SocialMediaHandler.GetSocialMedia)
	routers.PUT("/SocialMedias/:socialMediaId", SocialMediaHandler.UpdateSocialMedia)
	routers.DELETE("/SocialMedias/:socialMediaId", SocialMediaHandler.DeleteSocialMedia)
	// r.GET("/User", UserHandler.GetUserList)
	// r.GET("/User/:order_id", UserHandler.GetUserDetail)
	// r.DELETE("/User/:order_id", UserHandler.DeleteUser)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("=========== Server started ===========")
	r.Run(":1337")
}
