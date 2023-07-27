package domain

type BoqBody struct {
	Id                int     `json:"id" gorm:"column:id"`
	RunNum            string  `json:"run_num" gorm:"column:run_num"`
	ItemNo            string  `json:"item_no" gorm:"column:item_no"`
	ItemLevel         int     `json:"item_level" gorm:"column:item_level"`
	ItemDescription   string  `json:"item_description" gorm:"column:item_description"`
	ItemSpecification string  `json:"item_specification" gorm:"column:item_specification"`
	Qty               float32 `json:"qty" gorm:"column:qty"`
	Unit              string  `json:"unit" gorm:"column:unit"`
	Price             float64 `json:"price" gorm:"column:price"`
	Currency          string  `json:"currency" gorm:"column:currency"`
	Note              string  `json:"note" gorm:"column:note"`
}

type BoqBodyRequest struct {
	RunNum            string  `json:"run_num"`
	ItemNo            string  `json:"item_no"`
	ItemLevel         int     `json:"item_level"`
	ItemDescription   string  `json:"item_description"`
	ItemSpecification string  `json:"item_specification"`
	Qty               float32 `json:"qty"`
	Unit              string  `json:"unit"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	Note              string  `json:"note"`
}

type BoqBodyResponse struct {
	Id                int               `json:"id"`
	RunNum            string            `json:"run_num"`
	ItemNo            string            `json:"item_no"`
	ItemLevel         int               `json:"item_level"`
	ItemDescription   string            `json:"item_description"`
	ItemSpecification string            `json:"item_specification"`
	Qty               float32           `json:"qty"`
	Unit              string            `json:"unit"`
	Price             float64           `json:"price"`
	Currency          string            `json:"currency"`
	Note              string            `json:"note"`
	Children          []BoqBodyResponse `json:"children"`
	ParentId          int               `json:"parent_id"`
}

type BoqBodyResponseFinal struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []BoqBody `json:"data"`
}
