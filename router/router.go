package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/controller"
	"github.com/hyperjiang/gin-skeleton/middleware"
	"github.com/hyperjiang/gin-skeleton/model"
)

// Route makes the routing
func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	app.GET(
		"/", indexController.GetIndex,
	)

	auth := app.Group("/auth")
	authMiddleware := middleware.Auth()
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get("email")
			c.JSON(200, gin.H{
				"email": claims["email"],
				"name":  user.(*model.User).Name,
				"text":  "Hello World.",
			})
		})
	}

	userController := new(controller.UserController)
	app.GET(
		"/user/:id", userController.GetUser,
	).GET(
		"/signup", userController.SignupForm,
	).POST(
		"/signup", userController.Signup,
	).GET(
		"/login", userController.LoginForm,
	).POST(
		"/login", authMiddleware.LoginHandler,
	)

	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}

	authController := new(controller.AuthController)
	api.POST("/login", authMiddleware.LoginHandler)
	api.POST("/signup", authController.DoSignup)

	todoController := new(controller.TodoController)
	api.POST("/todo", todoController.StoreTodo)
	api.GET("/todo/:id", todoController.GetTodoById)
	api.PUT("/todo/:id", todoController.UpdateTodo)
	api.DELETE("/todo/:id", todoController.DeleteTodo)

	postController := new(controller.PostController)
	api.GET("/post", postController.GetPostsPaginate)
	api.POST("/post", postController.StorePost)
	api.GET("/post/:id", postController.GetTodoById)
	api.PUT("/post/:id", postController.UpdatePost)
	api.DELETE("/post/:id", postController.DeletePost)
}
