package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = config.GetCollection(config.DB, "products")

func GetAllProducts() string {
	return "Get all products"
}

func CreateProduct() (model.Product, error) {
	// validate := validator.New()
	// productSchema := model.Product

	// if err := c.BindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	// 	return
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// var product model.Product
	defer cancel()

	newProduct := model.Product{
		ProductName:  "Test product",
		ProductPrice: "100฿",
		Amount:       10,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := productCollection.InsertOne(ctx, newProduct)
	return newProduct, err
}
