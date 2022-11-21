package sxorm

import "time"

type (
	User struct {
		Id int `xorm:"not null pk autoincr INT(11) comment('id') <-"`
		// Id   int64
		Name string `xorm:"nvarchar(50) comment('姓名')"`
		Sex  int    `xorm:"INT(1)"`
		// CreateDate time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP <-"`
		CreateDate time.Time `xorm:"created"`
		UpdateTime time.Time `xorm:"updated"`
		// UpdateTime time.Time `xorm:"null default null TIMESTAMP <-"`
	}

	Account struct {
		Id      int64
		UserId  int64  `xorm:"index"`
		Account string `xorm:"VARCHAR(50)"`
	}
)
