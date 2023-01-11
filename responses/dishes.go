package responses

type DishRequestCreate struct {
	Name        string   `json:"name" bson:"name"`
	Ingredients []string `json:"ingredients" bson:"ingredients"`
}
