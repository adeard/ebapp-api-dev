package domain

type BoqBody struct {
	Id                int     `json:"id"`
	RunNum            string  `json:"run_num"`
	ItemNo            string  `json:"item_no"`
	ItemLevel         int     `json:"item_level"`
	ItemDescription   string  `json:"item_description"`
	ItemSpecification string  `json:"item_specification"`
	Qty               int     `json:"qty"`
	Unit              string  `json:"unit"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	Note              string  `json:"note"`
}

type BoqBodyRequest struct {
	RunNum            string  `json:"run_num"`
	ItemNo            string  `json:"item_no"`
	ItemLevel         int     `json:"item_level"`
	ItemDescription   string  `json:"item_description"`
	ItemSpecification string  `json:"item_specification"`
	Qty               int     `json:"qty"`
	Unit              string  `json:"unit"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	Note              string  `json:"note"`
}
