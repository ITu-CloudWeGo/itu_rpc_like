package model

type PidCount struct {
	Pid   int64 `gorm:"column:pid;not null" redis:"pid"`
	Count int64 `gorm:"column:count;default:0;not null" redis:"count"`
}

func (l *PidCount) PidCountTableName() string {
	return "likes"
}
