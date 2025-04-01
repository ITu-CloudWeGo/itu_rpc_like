package model

import (
	"gorm.io/plugin/optimisticlock"
)

type PidCount struct {
	Pid     int64 `gorm:"column:pid;not null" redis:"pid"`
	Count   int64 `gorm:"column:count;default:0;not null" redis:"count"`
	Version optimisticlock.Version
}

func (l *Likes) PidCountTableName() string {
	return "likes"
}
