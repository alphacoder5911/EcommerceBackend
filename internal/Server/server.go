// Http server
package server

import (
	"net/http"

	services "github.com/alphacoder5911/EcommerceBackend/internal/Services"
	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	Config *config.Config
	db     *gorm.DB
	logger zerolog.Logger
	authService  *services.AuthService
	userService *services.UserService
	productService *services.ProductService
}

func NewServer(config *config.Config, db *gorm.DB, logger zerolog.Logger) *Server {
	return &Server{
		Config: config,
		db:     db,
		logger: logger,
		authService: services.NewAuthService(db,config),
		userService: services.NewUserService(db),
		productService: services.NewProductService(db),
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())
	
	// Add routes

	router.GET("/health", s.healthCheck)

	api:=router.Group("/api/v1")
	{
		auth:=api.Group("/auth")
		{//nolint:gocritic // I need this for readability
			auth.POST("/register",s.Register)
			auth.POST("/login",s.Login)
			auth.POST("/refresh",s.RefreshToken)
			auth.POST("/logout",s.Logout)
		}

		protected:=api.Group("/prot")
		protected.Use(s.authMiddleware())
		{
			protected.GET("/profile",s.profile)
			protected.GET("/GetProfile",s.getProfile)
			protected.POST("/UpdateProfile/:id",s.UpdateProfile)
			protected.POST("/AddCategory",s.CreateCategory)
			protected.GET("/GetCategories",s.GeetCategories)
			protected.POST("/UpdateCategory/:id",s.UpdateCategory)
			protected.POST("/DeleteCategory/:id",s.DeleteCategory)
			protected.POST("/CreateProduct",s.CreateProduct)
			protected.GET("/GetProducts",s.getProducts)
			protected.GET("/GetProduct/:id",s.GetProduct)
			protected.POST("/UpdateProduct/:id",s.UpdateProduct)
			protected.POST("/DeleteProduct/:id",s.DeleteProduct)
		}
	}


	return router
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
