package dao

import (
	"errors"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	"gorm.io/gorm"
	"sync"
)

type PidCountDao struct {
	db *gorm.DB
}

type PidCountDaoImpl interface {
	GetLikesCount(pid int64) (int64, error)
	UpdateLikesCount(pid, count int64) error
	Insert(pidCount *model.PidCount) error
}

var (
	instancePidCountDAO *PidCountDao
	oncePidCountDAO     sync.Once
	noExists            int64 = -1
	failed              int64 = -2
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
func (dao *PidCountDao) GetLikesCount(pid int64) (int64, error) {
	var pidCount model.PidCount
	err := dao.db.Model(&model.PidCount{}).Where("pid = ?", pid).Update("count", pidCount).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return noExists, nil
		}
		return failed, err
	}
	return pidCount.Count, nil
}

func (dao *PidCountDao) UpdateLikesCount(pid, count int64) error {
	return dao.db.Model(&model.PidCount{}).Where("pid = ?", pid).Update("count", count).Error
}
