package domain

type PoPic struct {
	RunNum string `json:"run_num" gorm:"column:pekerjaan_no"`
	Id     int    `json:"id" gorm:"column:level"`
	Uid    string `json:"uid" gorm:"column:uid"`
	Name   string `json:"name" gorm:"column:name"`
	Email  string `json:"email" gorm:"column:email"`
	Role   string `json:"jabatan" gorm:"column:role"`
	Status string `json:"status" gorm:"column:status"`
}

type PoPicRequest struct {
	RunNum string `json:"run_num" gorm:"column:pekerjaan_no"`
	Id     int    `json:"id" gorm:"column:level"`
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"jabatan"`
	Status string `json:"status"`
}

type PoPicResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []PoPic `json:"data"`
}
