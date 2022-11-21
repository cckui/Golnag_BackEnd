package sxorm

import "time"

type TbInfra struct {
	PrimaryId  int64     `xorm:"not null pk autoincr INT(11)"`
	Data1      string    `xorm:"not null unique VARCHAR(50)"`
	Data2      string    `xorm:"not null VARCHAR(100)"`
	Data3      int       `xorm:"not null INTEGER"`
	Data4      bool      `xorm:"not null BOOL"`
	CreateDate time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
