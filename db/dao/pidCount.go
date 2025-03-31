package dao

import (
	"github.com/ITu-CloudWeGo/itu_rpc_like/db"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	"gorm.io/gorm"
	"sync"
)

type PidCountDao struct {
	db *gorm.DB
}

type PidCountDaoImpl interface {
	GetLikesCount(pid int64) int64
	UpdateLikesCount(pid, count int64) error
	Insert(pidCount *model.PidCount) error
}

var (
	instancePidCountDAO *PidCountDao
	oncePidCountDAO     sync.Once
)

func GetPidCountDao() PidCountDaoImpl {
	oncePidCountDAO.Do(func() {
		instancePidCountDAO = &PidCountDao{
			db: db.GetDB(),
		}
	})
	return instancePidCountDAO
}

func (dao *PidCountDao) Insert(pidCount *model.PidCount) error {
	return dao.db.Create(pidCount).Error
}
func (dao *PidCountDao) GetLikesCount(pid int64) int64 {
	var count int64
	dao.db.Model(&model.Likes{}).Where("pid = ?", pid).Count(&count)
	return count
}

func (dao *PidCountDao) UpdateLikesCount(pid, count int64) error {
	return dao.db.Model(&model.PidCount{}).Where("pid = ?", pid).Update("count", count).Error
}
