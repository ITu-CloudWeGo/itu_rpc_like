package model

import "fmt"

type Like struct {
	Pid int64 `gorm:"column:pid;uniqueIndex:idx_pid_uid;not null" redis:"pid"`
	Uid int64 `gorm:"column:uid;uniqueIndex:idx_pid_uid;not null" redis:"uid"`
}

func (l *Like) LikesTableName() string {
	return "likes"
}

func (l *Like) GetCacheKey() string {
	return fmt.Sprintf("%s:%d", l.LikesTableName())
}
