package domain

type PoBoqHeaderProgress struct {
	PekerjaanNo string  `json:"pekerjaan_no"`
	Po          string  `json:"po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"sum" gorm:"column:qty"`
	Unit        string  `json:"pounit" gorm:"column:unit"`
	Price       float32 `json:"price"`
	Currency    string  `json:"currency"`
	Order       string  `json:"order"`
	IsAddendum  int     `json:"is_addendum" gorm:"column:is_addendum"`
}

type PoBoqHeaderProgressResponse struct {
	Message string                `json:"message"`
	Status  int                   `json:"status"`
	Data    []PoBoqHeaderProgress `json:"data"`
}
