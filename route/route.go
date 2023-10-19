package route

import (
	"go-mini-project/config"
	"go-mini-project/controller"
	"go-mini-project/repository"
	"go-mini-project/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	// Clean Architecture
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository(db)
	bookService := usecase.NewBookUsecase(bookRepository)
	bookController := controller.NewBookController(bookService)

	//JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.JWT_KEY)))

	// Route to handler function
	e.POST("/users/signin", userController.SignIn)
	e.POST("/users", userController.SignUp)

	e.GET("/books", bookController.GetBooks)
	e.GET("/books/:id", bookController.GetBookById)
	r.POST("/books", bookController.InsertBook)
	r.PUT("/books/:id", bookController.UpdateBook)
	r.DELETE("/books/:id", bookController.DeleteBook)
}
