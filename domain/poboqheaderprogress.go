package domain

type PoBoqHeaderProgress struct {
	PekerjaanNo      string  `json:"pekerjaan_no"`
	Po               string  `json:"po"`
	Item             string  `json:"item"`
	Description      string  `json:"description"`
	Qty              float32 `json:"sum" gorm:"column:qty"`
	Unit             string  `json:"pounit" gorm:"column:unit"`
	Price            float64 `json:"price"`
	Currency         string  `json:"currency"`
	Order            string  `json:"order"`
	IsAddendum       int     `json:"is_addendum" gorm:"column:is_addendum"`
	Percentage       string  `json:"percentage" gorm:"column:percentage"`
	ActualPercentage string  `json:"actual_percentage" gorm:"column:actual_percentage"`
}

type PoBoqHeaderProgressResponse struct {
	Message string                `json:"message"`
	Status  int                   `json:"status"`
	Data    []PoBoqHeaderProgress `json:"data"`
}

type PoBoqHeaderProgressWithBodyServerSide struct {
	PoBoqHeaderProgress
	BoqBodyProgress []PoBoqBodyProgressResponse `json:"children"`
}
