package responses

type MenuRequestCreate struct {
	Name        string   `json:"name"`
	RawMaterial []string `json:"rawMaterial"`
}
