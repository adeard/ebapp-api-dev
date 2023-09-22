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

type Company struct {
	Bukrs string `json:"bukrs" gorm:"column:BUKRS"`
	Butxt string `json:"butxt" gorm:"column:BUTXT"`
}

type Plant struct {
	Werks string `json:"werks" gorm:"column:WERKS"`
	Name1 string `json:"plant" gorm:"column:NAME1"`
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

type AddonResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Company `json:"data"`
}

type AddonResponse2 struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Plant `json:"data"`
}
