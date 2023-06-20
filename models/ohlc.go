package models

type OHLC struct {
	Id     uint16  `json:"id"`
	Date   float32 `json:"date"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume float32 `json:"volume"`
}
