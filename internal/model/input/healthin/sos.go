package healthin

type SOSCreateInp struct {
	ElderUserId int    `json:"elderUserId" v:"required"`
	Longitude   string `json:"longitude" v:"required"`
	Latitude    string `json:"latitude" v:"required"`
	Notes       string `json:"notes" v:"required"`
}

type SOSCreateModel struct {
}
