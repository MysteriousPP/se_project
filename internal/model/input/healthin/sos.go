package healthin

import "github.com/gogf/gf/v2/os/gtime"

type SOSCreateInp struct {
	ElderUserId int         `json:"elderId" v:"required"`
	Time        *gtime.Time `json:"time"`
	Location    string      `json:"location"`
	Type        string      `json:"type"`
	Longitude   string      `json:"longitude"`
	Latitude    string      `json:"latitude"`
	Notes       string      `json:"notes"`
}

type SOSCreateModel struct {
	SosId   int         `json:"sosId"`
	ElderId int         `json:"elderId"`
	Time    *gtime.Time `json:"time"`
	Status  string      `json:"status"`
}
