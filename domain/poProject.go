package domain

type PoProject struct {
	Id          int     `json:"id"`
	PekerjaanNo string  `json:"pekerjaan_no"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"sum" gorm:"column:qty"`
	PoUnit      string  `json:"pounit" gorm:"column:unit"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Wbs         string  `json:"wbs"`
	Cera        string  `json:"cera"`
}

type PoProjectRequest struct {
	Id          int     `json:"id"`
	PekerjaanNo string  `json:"pekerjaan_no"`
	Po          string  `json:"po" gorm:"column:po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float32 `json:"sum" gorm:"column:qty"`
	PoUnit      string  `json:"pounit" gorm:"column:unit"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
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
