package dto

// Http accepted representation of add locker request
type AddLockerHttpInboundDto struct {
	ClientId string `json:"clientid"`
}

// Http accepted representation of add locker response
type AddLockerHttpOutboundDto struct {
	LockerId string `json:"lockerid"`
}

// Http accepted representation of add item request
type AddItemHttpInboundDto struct {
	ClientId string `json:"clientid"`
	LockerId string `json:"lockerid"`
	SecretId string `json:"secretid"`
	Content  string `json:"content"`
}

// Http accepted representation of add item response
type AddItemHttpOutboundDto struct {
	Status bool `json:"status"`
}

// Http accepted representation of remove item request
type RemoveItemHttpInboundDto struct {
	ClientId string `json:"clientid"`
	LockerId string `json:"lockerid"`
	SecretId string `json:"secretid"`
}

// Http accepted representation of remove item response
type RemoveItemHttpOutboundDto struct {
	Status bool `json:"status"`
}

// Http accepted representation of get item request
type GetItemHttpInboundDto struct {
	ClientId string `json:"clientid"`
	LockerId string `json:"lockerid"`
	SecretId string `json:"secretid"`
}

// Http accepted representation of get item response
type GetItemHttpOutboundDto struct {
	Content string `json:"content"`
}
