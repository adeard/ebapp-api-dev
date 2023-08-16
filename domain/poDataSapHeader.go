package domain

import (
	"encoding/xml"
	"fmt"
	"time"
)

type entry struct {
	XMLName xml.Name `xml:"entry"`
	Content content  `xml:"content"`
}

type content struct {
	XMLName    xml.Name   `xml:"content"`
	Properties properties `xml:"properties"`
}

type properties struct {
	XMLName   xml.Name  `xml:"properties"`
	PoDate    time.Time `xml:"d:DocDate"`
	Pt        string    `xml:"d:CompCode"`
	UnitUsaha string    `xml:"d:Plant"`
	Vendor    string    `xml:"d:Vendor"`
}

func ParseXML(xmlData []byte) (PoDataSapHeader, error) {
	var entryData entry
	err := xml.Unmarshal(xmlData, &entryData)
	if err != nil {
		return PoDataSapHeader{}, err
	}

	poHeader := PoDataSapHeader{
		PoDate:    entryData.Content.Properties.PoDate,
		Pt:        entryData.Content.Properties.Pt,
		UnitUsaha: entryData.Content.Properties.UnitUsaha,
		Vendor:    entryData.Content.Properties.Vendor,
	}

	fmt.Println("Data yang di ambil :", poHeader)

	return poHeader, nil
}

type PoDataSapHeader struct {
	PoDate    time.Time
	Pt        string
	UnitUsaha string
	Vendor    string
}

type PoDataSapHeaderResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []PoDataSapHeader `json:"data"`
}
