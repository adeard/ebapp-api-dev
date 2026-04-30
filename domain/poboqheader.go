package domain

type PoBoqHeader struct {
	PekerjaanNo string  `json:"pekerjaan_no"`
	Po          string  `json:"po"`
	Item        string  `json:"item"`
	Description string  `json:"description"`
	Qty         float64 `json:"sum" gorm:"column:qty"`
	Unit        string  `json:"pounit" gorm:"column:unit"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Order       string  `json:"order"`
	IsAddendum  int     `json:"is_addendum" gorm:"column:is_addendum"`
	ActualPrice float64 `json:"actual_price"`
}

type PoBoqHeaderResponse struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    []PoBoqHeader `json:"data"`
}

type PoBoqHeaderWithBody struct {
	PoBoqHeader
	BoqBody []PoBoqBodyResponse `json:"children"`
}

type PoBoqHeaderWithBodyServerSide struct {
	PoBoqHeader
	BoqBody []PoBoqBodyServerSideResponse `json:"children"`
}

type PoBoqHeaderFilterRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type PoBoqHeaderFilterRequestServerSide struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Item     string `form:"item"`
	Desc     string `form:"desc"`
}
