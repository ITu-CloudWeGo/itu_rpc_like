package module

type Likes struct {
	Pid uint64 `gorm:"column:pid;primaryKey" json:"pid"`
	Uid uint64 `gorm:"column:uid;not null" json:"uid"`
}

type Tabler interface {
	TableName() string
}

func (l *Likes) TableName() string {
	return "likes"
}
