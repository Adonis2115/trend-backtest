package models

type OHLC struct {
	ID     uint16
	date   string
	open   float32
	high   float32
	low    float32
	close  float32
	volume uint64
}
