package models

type Router struct {
	Name  string `json:"name"`
	Datas []Data `json:"data"`
}

type Data struct {
	Time   int `json:"time"`
	Uptime int `json:"uptime"`
}
