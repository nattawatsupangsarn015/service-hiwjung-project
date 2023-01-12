package responses

type DishRequestCreate struct {
	Name        string                  `json:"name" bson:"name"`
	Ingredients []DishRequestIngredient `json:"ingredients" bson:"ingredients"`
}

type DishRequestIngredient struct {
	Value  string  `json:"value" bson:"value"`
	Unit   string  `json:"unit" bson:"unit"`
	Amount float64 `json:"amount" bson:"amount"`
}
