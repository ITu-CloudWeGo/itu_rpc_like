package model

import "fmt"

type Likes struct {
	ID  int64 `gorm:"column:id;primaryKey;autoIncrement" redis:"id"`
	Pid int64 `gorm:"column:pid;uniqueIndex:idx_pid_uid;not null" redis:"pid"`
	Uid int64 `gorm:"column:uid;uniqueIndex:idx_pid_uid;not null" redis:"uid"`
}

func (l *Likes) TableName() string {
	return "likes"
}

func (l *Likes) GetCacheKey() string {
	return fmt.Sprintf("%s:%d", l.TableName(), l.ID)
}
