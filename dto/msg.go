package dto

type AddLockerRequest struct {
	Client   string `json:"clientid"`
	Password string `json:"password"`
}

type AddLockerResponse struct {
	LockerId string `json:"lockerid"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
