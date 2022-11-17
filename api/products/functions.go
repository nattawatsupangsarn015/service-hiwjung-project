package products

import (
	"example/service-hiwjung-project/controller"

	"github.com/go-playground/validator/v10"
)

func GetAllProducts() string {
	return controller.GetAllProducts()
}

func CreateProducts(product interface{}) (interface{}, error) {
	validate := validator.New()
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil {
		return nil, nil
	}

	return product, nil
}
