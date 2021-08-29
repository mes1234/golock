package dto

// Http accepted representation of add locker request
type AddLockerHttpInboundDto struct {
	Client   string `json:"clientid"`
	Password string `json:"password"`
}

// Http accepted representation of add locker response
type AddLockerHttpOutboundDto struct {
	LockerId string `json:"lockerid"`
}
