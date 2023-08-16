package domain

import (
	"time"
)

type PoDataSapHeader struct {
	PoDate    time.Time `json:"date"`       //<d:DocDate>2018-02-06T00:00:00</d:DocDate>
	Pt        string    `json:"pt"`         //<d:CompCode>3500</d:CompCode>
	UnitUsaha string    `json:"unit_usaha"` //<d:Plant>3520</d:Plant>
	Vendor    string    `json:"vendor"`     //<d:Vendor>100016</d:Vendor>
}

type PoDataSapHeaderResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []PoDataSapHeader `json:"data"`
}
