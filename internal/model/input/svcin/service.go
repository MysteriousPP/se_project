package svcin

type ServiceOrderCreateInp struct {
	ServiceId   int `json:"serviceId" v:"required"`
	ElderUserId int `json:"elderUserId" v:"required"`
	BookedBy    int `json:"bookedBy" v:"required"`
}

type ServiceOrderCreateModel struct {
	BookingId int `json:"bookingId"`
	Status    int `json:"status"`
}
