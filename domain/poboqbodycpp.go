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

type PoBoqBodyCppProgress struct {
	Status            bool    `json:"status_cpp" gorm:"column:status_cpp"`
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
	Note              string  `json:"note_cpp" gorm:"column:note_cpp"`
	StartDate         string  `json:"start_date" gorm:"column:start_date;default:NULL"`
	EndDate           string  `json:"end_date" gorm:"column:end_date;default:NULL"`
	StartDateActual   string  `json:"start_date_actual" gorm:"column:start_date_actual;default:NULL"`
	EndDateActual     string  `json:"end_date_actual" gorm:"column:end_date_actual;default:NULL"`
	PreviousVolume    float64 `json:"prev_volume" gorm:"column:previous_volume;default:NULL"`
	CurrentVolume     float64 `json:"current_volume" gorm:"column:current_volume;default:NULL"`
	RunNumCpp         string  `json:"run_num_cpp" gorm:"column:run_num_cpp"`
}

type PoBoqBodyCppProgressRequest struct {
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

type PoBoqBodyCppProgressResponse struct {
	Status            bool                           `json:"status_cpp" gorm:"column:status_cpp"`
	Id                int                            `json:"id" gorm:"column:main_id"`
	ParentId          int                            `json:"parent_id" gorm:"parent_id"`
	RunNum            string                         `json:"run_num" gorm:"column:run_num"`
	Order             string                         `json:"order" gorm:"column:order"`
	ItemNo            string                         `json:"item_no" gorm:"column:item_no"`
	ItemLevel         int                            `json:"item_level" gorm:"column:item_level"`
	ItemDescription   string                         `json:"item_description" gorm:"column:item_description"`
	ItemSpecification string                         `json:"item_specification" gorm:"column:item_specification"`
	Qty               float32                        `json:"qty" gorm:"column:qty"`
	Unit              string                         `json:"unit" gorm:"column:unit"`
	Price             float64                        `json:"price" gorm:"column:price"`
	Currency          string                         `json:"currency" gorm:"column:currency"`
	Note              string                         `json:"note_cpp" gorm:"column:note_cpp"`
	StartDate         string                         `json:"start_date" gorm:"column:start_date;default:NULL"`
	EndDate           string                         `json:"end_date" gorm:"column:end_date;default:NULL"`
	StartDateActual   string                         `json:"start_date_actual" gorm:"column:start_date_actual;default:NULL"`
	EndDateActual     string                         `json:"end_date_actual" gorm:"column:end_date_actual;default:NULL"`
	PreviousVolume    float64                        `json:"prev_volume" gorm:"column:previous_volume;default:NULL"`
	CurrentVolume     float64                        `json:"current_volume" gorm:"column:current_volume;default:NULL"`
	RunNumCpp         string                         `json:"run_num_cpp" gorm:"column:run_num_cpp"`
	Children          []PoBoqBodyCppProgressResponse `json:"children"`
}

type PoBoqBodyCppProgressResponseFinal struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    []PoBoqBodyCppProgress `json:"data"`
}
