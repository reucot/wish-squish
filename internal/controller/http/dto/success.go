package dto

type Success struct {
	Ok   string      `json:"ok"`
	Data interface{} `json:"data"`
}
