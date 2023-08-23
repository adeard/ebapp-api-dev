package podatasapheader

import (
	"ebapp-api-dev/domain"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type Repository interface {
	CheckTitle(id string) ([]domain.PoDataSapHeaderTitle, error)
	CheckArea(id string) (domain.DataMasterPlant, error)
	CheckWbs(id string) ([]domain.ReadWbs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CheckTitle(id string) ([]domain.PoDataSapHeaderTitle, error) {
	var poProject []domain.PoDataSapHeaderTitle

	//Setting
	username := "SIT_GUNAWAN"
	password := "acer620"

	// FIRST URL
	xmlURL := fmt.Sprintf(`http://qaecc.hec.indofood.co.id:8020/sap/opu/odata/sap/ZMGW_GET_DATA_PO_SRV/etHeaderSet('` + id + `')?$expand=etPoHeaderSet,NavPoItemSet`)

	client := &http.Client{}
	req, err := http.NewRequest("GET", xmlURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Charset", "utf-8")
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

	poHeaderTitle, err := domain.ParseXMLTitle(result)
	if err != nil {
		return nil, err
	}

	poProject = append(poProject, poHeaderTitle)

	return poProject, nil
}

func (r *repository) CheckWbs(id string) ([]domain.ReadWbs, error) {
	var wbsPo []domain.ReadWbs

	//Setting
	username := "SIT_GUNAWAN"
	password := "acer620"

	xmlURLWBS := fmt.Sprintf(`http://qaecc.hec.indofood.co.id:8020/sap/opu/odata/sap/ZMGW_GET_DATA_PO_SRV/etHeaderSet('` + id + `')?$expand=NavPoAccountSet`)

	client := &http.Client{}
	req, err := http.NewRequest("GET", xmlURLWBS, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Charset", "utf-8")
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

	poHeaderTitle2, err := domain.ParseXMLWbs(result)
	if err != nil {
		return nil, err
	}

	wbsPoResult := append(wbsPo, poHeaderTitle2)

	return wbsPoResult, nil
}

func (r *repository) CheckArea(id string) (domain.DataMasterPlant, error) {
	var DataMasterPlant domain.DataMasterPlant
	err := r.db.Table("MasterPlantArea").Where("WERKS = ?", id).First(&DataMasterPlant).Error
	return DataMasterPlant, err
}
