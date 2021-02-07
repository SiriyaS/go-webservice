package server

import (
	"backend-shortcourse/go-webservice/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend-shortcourse/go-webservice/docs"
)

// @title My Shop
// @version Backend-Shortcourse
// @description This is a My Shop API document from Backend-Shortcourse workshop

// @host localhost:4000
// @scheme http

// @tag.name Product
// @tag.description create read update delete product
// @tag.name Ping
// @tag.description check server status

func NewRouter() *gin.Engine {
	router := gin.New()

	url := ginSwagger.URL("http://localhost:4000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/ping", controller.Ping)

	productController := controller.ProductController{}
	router.GET("/products", productController.ReadAll)
	router.GET("/product", productController.Read) // pass product_id through query string
	router.POST("/product", productController.Add)
	router.PUT("/update", productController.Update)    // pass product_id through query string
	router.DELETE("/delete", productController.Delete) // pass product_id through query string

	return router
}
