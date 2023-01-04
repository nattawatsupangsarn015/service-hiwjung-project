package responses

type IngredientRequestCreate struct {
	Name       string   `json:"name"`
	Ingredient []string `json:"Ingredient"`
}
