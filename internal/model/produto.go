package model

type Produto struct {
	ID    uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome  string  `json:"name"`
	Preco float64 `json:"price"`
}
