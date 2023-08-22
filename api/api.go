package api

import (
	_ "github.com/user/api/docs"
	"github.com/user/api/handler"
	"github.com/user/config"
	"github.com/user/pkg/logger"
	"github.com/user/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")

	// Login Api
	r.POST("/login", handler.Login)

	// Register Api
	r.POST("/register", handler.Register)

	// User Api
	v1.Use(handler.AuthMiddleware())
	v1.POST("/user", handler.CreateUser)
	v1.GET("/user/:id", handler.GetByIdUser)
	v1.GET("/user", handler.GetListUser)
	v1.PUT("/user/:id", handler.UpdateUser)
	v1.DELETE("/user/:id", handler.DeleteUser)

	//PhoneNumber Api
	v1.POST("/phone_number", handler.CreatePhoneNumber)
	v1.GET("/phone_number/:id", handler.GetByIdPhoneNumber)
	v1.GET("/phone_number", handler.GetListPhoneNumber)
	v1.PUT("/phone_number/:id", handler.UpdatePhoneNumber)
	v1.DELETE("/phone_number/:id", handler.DeletePhoneNumber)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
