package domain

import (
	"time"
)

type ParEntries struct {
	Id            int       `json:"id"`
	TabId         int       `json:"tab_id"`
	Sdesc         string    `json:"sdesc"`
	Ldesc         string    `json:"ldesc"`
	van1          int       `json:"van1"`
	van2          int       `json:"van2"`
	van3          int       `json:"van3"`
	van4          int       `json:"van4"`
	vac1          int       `json:"vac1"`
	vac2          int       `json:"vac2"`
	vac3          int       `json:"vac3"`
	vac4          int       `json:"vac4"`
	CreateDate    time.Time `json:"createdate"`
	CreatedBy     string    `json:"createdby"`
	LastUpdated   time.Time `json:"lastupdated"`
	LastUpdatedBy string    `json:"lastupdated"`
}

type ParEntriesRequest struct {
	Id            int       `json:"id"`
	TabId         int       `json:"tab_id"`
	Sdesc         string    `json:"sdesc"`
	Ldesc         string    `json:"ldesc"`
	van1          int       `json:"van1"`
	van2          int       `json:"van2"`
	van3          int       `json:"van3"`
	van4          int       `json:"van4"`
	vac1          int       `json:"vac1"`
	vac2          int       `json:"vac2"`
	vac3          int       `json:"vac3"`
	vac4          int       `json:"vac4"`
	CreateDate    time.Time `json:"createdate"`
	CreatedBy     string    `json:"createdby"`
	LastUpdated   time.Time `json:"lastupdated"`
	LastUpdatedBy string    `json:"lastupdated"`
}

type ParEntriesResponse struct {
	Data    []ParEntries `json:"data"`
	Status  int          `json:"status"`
	Message string       `json:"message"`
}
