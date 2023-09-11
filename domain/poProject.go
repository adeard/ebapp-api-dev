package domain

type PoProject struct {
	Id          int     `json:"id"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"qty"`
	Price       float64 `json:"price"`
	Wbs         string  `json:"wbs"`
	Cera        string  `json:"cera"`
}

type PoProjectRequest struct {
	Id          int     `json:"id"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"qty"`
	Price       float64 `json:"price"`
	Wbs         string  `json:"wbs"`
	Cera        string  `json:"cera"`
}

type PoProjectResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []PoProject `json:"data"`
}

type PoProjectResponse2 struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
