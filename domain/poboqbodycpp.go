package domain

type PoBoqBodyCpp struct {
	Id                int     `json:"id" gorm:"column:main_id"`
	ParentId          int     `json:"parent_id" gorm:"parent_id"`
	RunNum            string  `json:"run_num" gorm:"column:run_num"`
	Order             string  `json:"order" gorm:"column:order"`
	ItemNo            string  `json:"item_no" gorm:"column:item_no"`
	ItemLevel         int     `json:"item_level" gorm:"column:item_level"`
	ItemDescription   string  `json:"item_description" gorm:"column:item_description"`
	ItemSpecification string  `json:"item_specification" gorm:"column:item_specification"`
	Qty               float32 `json:"qty" gorm:"column:qty"`
	Status            bool    `json:"status" gorm:"column:status"`
	Note              string  `json:"note" gorm:"column:note"`
}

type PoBoqBodyCppRequest struct {
	Id                int     `json:"id" gorm:"column:main_id"`
	ParentId          int     `json:"parent_id"`
	RunNum            string  `json:"run_num"`
	Order             string  `json:"order"`
	ItemNo            string  `json:"item_no"`
	ItemLevel         int     `json:"item_level"`
	ItemDescription   string  `json:"item_description"`
	ItemSpecification string  `json:"item_specification"`
	Qty               float32 `json:"qty"`
	Status            bool    `json:"status"`
	Note              string  `json:"note"`
}

type PoBoqBodyCppResponse struct {
	Id                int                    `json:"id"`
	ParentId          int                    `json:"parent_id"`
	RunNum            string                 `json:"run_num"`
	Order             string                 `json:"order"`
	ItemNo            string                 `json:"item_no"`
	ItemLevel         int                    `json:"item_level"`
	ItemDescription   string                 `json:"item_description"`
	ItemSpecification string                 `json:"item_specification"`
	Qty               float32                `json:"qty"`
	Status            bool                   `json:"status"`
	Note              string                 `json:"note"`
	Children          []PoBoqBodyCppResponse `json:"children"`
}

type PoBoqBodyCppResponseFinal struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []PoBoqBodyCpp `json:"data"`
}
