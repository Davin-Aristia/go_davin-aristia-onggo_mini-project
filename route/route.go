package route

import (
	"go-mini-project/config"
	"go-mini-project/controller"
	"go-mini-project/repository"
	"go-mini-project/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	openai "github.com/sashabaranov/go-openai"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Clean Architecture
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository(db)
	bookService := usecase.NewBookUsecase(bookRepository)
	bookController := controller.NewBookController(bookService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	
	salesRepository := repository.NewSalesRepository(db)
	salesService := usecase.NewSalesUsecase(salesRepository, bookRepository)
	salesController := controller.NewSalesController(salesService)
	
	vendorRepository := repository.NewVendorRepository(db)
	vendorService := usecase.NewVendorUsecase(vendorRepository)
	vendorController := controller.NewVendorController(vendorService)
	
	purchaseRepository := repository.NewPurchaseRepository(db)
	purchaseService := usecase.NewPurchaseUsecase(purchaseRepository, bookRepository)
	purchaseController := controller.NewPurchaseController(purchaseService)
	
	client := openai.NewClient(config.CHATBOT_KEY)
	chatbotRepository := repository.NewChatbotRepository(client)
	chatbotService := usecase.NewChatbotUsecase(chatbotRepository, categoryRepository, bookRepository)
	chatbotController := controller.NewChatbotController(chatbotService)

	// Route to handler function
	// auth
	e.POST("/users/signin", userController.SignIn)
	e.POST("/users", userController.SignUp)

	// books
	books := e.Group("/books")
	books.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/books", bookController.GetBooks)
	e.GET("/books/:id", bookController.GetBookById)
	books.POST("", bookController.InsertBook)
	books.PUT("/:id", bookController.UpdateBook)
	books.DELETE("/:id", bookController.DeleteBook)

	// categories
	categories := e.Group("/categories")
	categories.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/categories", categoryController.GetCategories)
	e.GET("/categories/:id", categoryController.GetCategoryById)
	categories.POST("", categoryController.InsertCategory)
	categories.PUT("/:id", categoryController.UpdateCategory)
	categories.DELETE("/:id", categoryController.DeleteCategory)
	
	// sales
	sales := e.Group("/sales")
	sales.Use(middleware.JWT([]byte(config.JWT_KEY)))
	sales.POST("/checkout", salesController.Checkout)
	sales.GET("/:id", salesController.GetSalesById)
	sales.GET("", salesController.GetSales)
	
	// vendors
	vendors := e.Group("/vendors")
	vendors.Use(middleware.JWT([]byte(config.JWT_KEY)))
	vendors.GET("", vendorController.GetVendors)
	vendors.GET("/:id", vendorController.GetVendorById)
	vendors.POST("", vendorController.InsertVendor)
	vendors.PUT("/:id", vendorController.UpdateVendor)
	vendors.DELETE("/:id", vendorController.DeleteVendor)
	
	// purchases
	purchases := e.Group("/purchases")
	purchases.Use(middleware.JWT([]byte(config.JWT_KEY)))
	purchases.POST("", purchaseController.CreatePurchase)
	purchases.GET("/:id", purchaseController.GetPurchaseById)
	purchases.GET("", purchaseController.GetPurchase)

	// chatbots
	chatbots := e.Group("/chatbots")
	chatbots.Use(middleware.JWT([]byte(config.JWT_KEY)))
	chatbots.POST("/book-recommendation", chatbotController.BookRecommendation)
}
