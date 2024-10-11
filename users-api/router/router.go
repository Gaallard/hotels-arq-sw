package router

import (
	"time"
	users "users-api/users-api/controller"
	"users-api/users-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.Use(middleware.CORSMiddleware())
	engine.POST("/users/login", users.Login)
	engine.POST("/users/register", users.RegisterUser)

	//authorized := engine.Group("")
	//authorized.Use(middleware.AuthMiddleware())
	//{
	//	authorized.POST("/inscriptions/:idCourse", inscriptions.UserInscription)
	//	authorized.GET("/courses/:idUser", courses.GetCoursesByUser)
	//	authorized.GET("/courses", courses.GetCourses)
	//	authorized.GET("/course/:idCourse", courses.GetCourseById)
	//
	//	authorized.GET("/comments/:idCourse", comments.GetCommentsByCourse)
	//	authorized.POST("/comments", comments.CreateComment)
	//	authorized.DELETE("/comments/:idComment", comments.DeleteComment)
	//	admin := authorized.Group("")
	//	admin.Use(middleware.AuthMiddlewareAdmin())
	//	{
	//		admin.POST("/courses", courses.CreateCourse)
	//		admin.PUT("/courses", courses.UpdateCourse)
	//		admin.DELETE("/courses/delete/:idCourse", courses.DeleteCourse)
	//	}
	//}
}

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
