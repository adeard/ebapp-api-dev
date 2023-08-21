package domain

import (
	"encoding/xml"
)

type PoDataSapHeaderTitle struct {
	PoNumber  string   `xml:"link>inline>feed>entry>content>properties>PoNumber"`
	CompCode  string   `xml:"link>inline>feed>entry>content>properties>CompCode"`
	DocDate   string   `xml:"link>inline>feed>entry>content>properties>DocDate"`
	Vendor    string   `xml:"link>inline>feed>entry>content>properties>Vendor"`
	Plant     string   `xml:"link>inline>feed>entry>content>properties>Plant"`
	Currency  string   `xml:"link>inline>feed>entry>content>properties>Currency"`
	CreatedBy string   `xml:"link>inline>feed>entry>content>properties>CreatedBy"`
	PoItem    []PoItem `xml:"link>inline>feed>entry>content>properties>PoItem"`
}

type PoItem struct {
	PoItem string `xml:"link>inline>feed>entry>content>properties>PoItem"`
}

func ParseXMLTitle(xmlData []byte) (PoDataSapHeaderTitle, error) {
	var parsedData PoDataSapHeaderTitle

	err := xml.Unmarshal(xmlData, &parsedData)
	if err != nil {
		return PoDataSapHeaderTitle{}, err
	}

	return parsedData, nil
}

type PoDataSapHeaderTitleResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    []PoDataSapHeaderTitle `json:"data"`
}

type DataMasterPlant struct {
	Werks string `json:"werks" gorm:"column:WERKS"`
	Area  string `json:"area" gorm:"column:AREA"`
	Area1 string `json:"area1" gorm:"column:AREA1"`
}

type DataMasterPlantResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []DataMasterPlant `json:"data"`
}
