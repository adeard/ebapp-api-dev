package domain

type PoProjectAddendum struct {
	Id          int     `json:"id"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"qty"`
	Price       float64 `json:"price"`
	Wbs         string  `json:"wbs"`
	Cera        string  `json:"cera"`
}

type PoProjectAddendumRequest struct {
	Id          int     `json:"id"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"qty"`
	Price       float64 `json:"price"`
	Wbs         string  `json:"wbs"`
	Cera        string  `json:"cera"`
}

type PoProjectAddendumResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    []PoProjectAddendum `json:"data"`
}
