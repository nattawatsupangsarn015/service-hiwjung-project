package responses

type MenuRequestCreate struct {
	Name       string   `json:"name"`
	Ingredient []string `json:"Ingredient"`
}
