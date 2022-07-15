package dto

type Error struct {
	Code    int         `json:"error_code"`
	Message string      `json:"error_msg"`
	Data    interface{} `json:"data"`
}
