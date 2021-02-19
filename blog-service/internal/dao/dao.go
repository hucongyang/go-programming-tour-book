package dao

import "github.com/jinzhu/gorm"

// 在dao层进行了数据访问对象的封装，并对业务所需的字段进行了处理

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
