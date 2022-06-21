package models

type IpList struct {
	Id   int64
	Ip   string   `xorm:"ip"`
	Info []Source `xorm:"info"`
}

type Source struct {
	Desc   string
	Source string
}
