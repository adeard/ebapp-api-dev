package domain

type PoBoqBody struct {
	Id                int     `json:"id" gorm:"column:main_id"`
	ParentId          int     `json:"parent_id" gorm:"parent_id"`
	RunNum            string  `json:"run_num" gorm:"column:run_num"`
	Order             string  `json:"order" gorm:"column:order"`
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

type PoBoqBodyRequest struct {
	Id                int     `json:"id" gorm:"column:main_id"`
	ParentId          int     `json:"parent_id"`
	RunNum            string  `json:"run_num"`
	Order             string  `json:"order"`
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

type PoBoqBodyResponse struct {
	Id                int                 `json:"id"`
	ParentId          int                 `json:"parent_id"`
	RunNum            string              `json:"run_num"`
	Order             string              `json:"order"`
	ItemNo            string              `json:"item_no"`
	ItemLevel         int                 `json:"item_level"`
	ItemDescription   string              `json:"item_description"`
	ItemSpecification string              `json:"item_specification"`
	Qty               float32             `json:"qty"`
	Unit              string              `json:"unit"`
	Price             float64             `json:"price"`
	Currency          string              `json:"currency"`
	Note              string              `json:"note"`
	Children          []PoBoqBodyResponse `json:"children"`
}

type PoBoqBodyResponseFinal struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []PoBoqBody `json:"data"`
}
