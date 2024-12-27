package main

import (
	"github.com/example/layeredArchitectureWithGo/internal/infrastructure/memory"
	"github.com/example/layeredArchitectureWithGo/internal/interfaces/handlers"
	"github.com/example/layeredArchitectureWithGo/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 依存関係を注入
	workRepo := memory.NewWorkRepository()
	workUseCase := usecase.NewWorkUseCase(workRepo)
	workHandler := handlers.NewWorkHandler(workUseCase)

	// ルーティングを設定
	api := e.Group("/api")
	{
		works := api.Group("/works")
		works.GET("", workHandler.GetAll)
		works.GET("/:id", workHandler.GetByID)
		works.POST("", workHandler.Create)
		works.PUT("/:id", workHandler.Update)
		works.DELETE("/:id", workHandler.Delete)
	}

	// サーバーを起動
	e.Logger.Fatal(e.Start(":8080"))
}
