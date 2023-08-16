package podatasapheader

import (
	"ebapp-api-dev/domain"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPo(po string) ([]domain.PoDataSapHeader, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// func (r *repository) FindByPo(po string) ([]domain.PoDataSapHeader, error) {
// 	var poProject []domain.PoDataSapHeader

// 	xmlURL := fmt.Sprintf("http://qaecc.hec.indofood.co.id:8020/sap/opu/odata/sap/ZMGW_GET_DATA_PO_SRV/etHeaderSet('%s')?$expand=etPoHeaderSet", po)

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", xmlURL, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Set Basic Authentication header
// 	username := "SIT_GUNAWAN"
// 	password := "acer620"

// 	req.Header.Set("Content-Type", "Application/xml")
// 	req.SetBasicAuth(username, password)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	result, _ := ioutil.ReadAll(resp.Body)
// 	poHeader, err := domain.ParseXML(result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Id")
// 	fmt.Println("PoDate:", poHeader.PoDate)
// 	fmt.Println("Pt:", poHeader.Pt)
// 	fmt.Println("UnitUsaha:", poHeader.UnitUsaha)
// 	fmt.Println("Vendor:", poHeader.Vendor)

// 	return poProject, nil
// }

func (r *repository) FindByPo(po string) ([]domain.PoDataSapHeader, error) {
	var poProject []domain.PoDataSapHeader

	xmlURL := fmt.Sprintf("http://qaecc.hec.indofood.co.id:8020/sap/opu/odata/sap/ZMGW_GET_DATA_PO_SRV/etHeaderSet('%s')?$expand=etPoHeaderSet", po)

	client := &http.Client{}
	req, err := http.NewRequest("GET", xmlURL, nil)
	if err != nil {
		return nil, err
	}

	// Set Basic Authentication header
	username := "SIT_GUNAWAN"
	password := "acer620"

	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	poHeader, err := domain.ParseXML(result)
	if err != nil {
		return nil, err
	}

	poProject = append(poProject, poHeader) // Append the parsed header to the slice

	return poProject, nil
}
