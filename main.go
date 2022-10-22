package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nasubibocchi/go_rest_api_try/handler"
	infra "github.com/nasubibocchi/go_rest_api_try/infrastructure"
	"github.com/nasubibocchi/go_rest_api_try/service"
)

func main() {
	loadEnv()

	engine := infra.DBInit()
	factory := service.NewService(engine)

	defer func() {
		log.Println("---engine closed")
		engine.Close()
	}()

	r := gin.Default()
	r.Use(service.ServiceFactoryModdleware(factory))

	routes := r.Group("/v1")
	{
		routes.POST("/books", handler.CreateBook)
		routes.GET("/books", handler.GetAll)
		routes.GET("/books/:id", handler.GetById)
		routes.PUT("/books/:id", handler.Update)
		routes.DELETE("/books/:id", handler.Delete)
	}

	r.Run(":8080")
	fmt.Println("----local server port 8080 started")
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
