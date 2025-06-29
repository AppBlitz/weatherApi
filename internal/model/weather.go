package model

type Weather struct {
	Latitude       float64 `json:"latitude"`
	Address        string  `json:"address"`
	ResolveAddress string  `json:"resolvedAddress"`
	Description    string  `json:"description"`
	TimeZone       string  `json:"timezone"`
	Longitude      string  `json:"longitude"`
}
