package domain

type PoBoqHeaderCpp struct {
	PekerjaanNo         string  `json:"pekerjaan_no"`
	PekerjaanNoProgress string  `json:"pekerjaan_no_progress"`
	Po                  string  `json:"po"`
	Item                string  `json:"item"`
	Description         string  `json:"description"`
	Qty                 float32 `json:"sum" gorm:"column:qty"`
	Order               string  `json:"order"`
}

type PoBoqHeaderCppResponse struct {
	Message string           `json:"message"`
	Status  int              `json:"status"`
	Data    []PoBoqHeaderCpp `json:"data"`
}
