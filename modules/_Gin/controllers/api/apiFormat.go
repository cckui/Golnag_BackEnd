package api

const app_ver string = "1.0.0"

type infoJSON struct {
	Version  string
	SysTime  string
	Web_Port string
	Web_Mode string
	WS_IP    string
	DB_IP    string
}

type sample struct {
	Status uint16 `json:"Status"`
	Data   string `json:"Data"`
	Time   string `json:"Time"`
}
