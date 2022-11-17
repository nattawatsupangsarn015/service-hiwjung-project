package responses

type ProductRequestCreate struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Amount int    `json:"amount"`
}
