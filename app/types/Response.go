package types

// swagger:model
type Response struct {
	PayLoad interface{} `json:"payload"`
	Status  int    `json:"status"`
}
