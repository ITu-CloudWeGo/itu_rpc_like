package module

type Likes struct {
	Pid int64 `gorm:"column:pid;primaryKey" json:"pid"`
	Uid int64 `gorm:"column:uid;not null" json:"uid"`
}

type PUTable interface {
	PUTableName() string
}

func (l *Likes) PUTableName() string {
	return "likes"
}
