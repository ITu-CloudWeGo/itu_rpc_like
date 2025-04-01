package dao

import (
	"errors"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	"gorm.io/gorm"
	"sync"
)

type LikesDao struct {
	db *gorm.DB
}

type LikesDaoImpl interface {
	Insert(tags *model.Like) error
	DelLikes(pid, uid int64) error
	IsLiked(pid, uid int64) (bool, error)
}

var (
	instanceLikesDAO *LikesDao
	onceLikesDao     sync.Once
)

func GetLikesDao() LikesDaoImpl {
	onceLikesDao.Do(func() {
		instanceLikesDAO = &LikesDao{
			db: db.GetDB(),
		}
	})
	return instanceLikesDAO
}

func (dao *LikesDao) Insert(likes *model.Like) error {
	return dao.db.Create(likes).Error
}

func (dao *LikesDao) DelLikes(pid, uid int64) error {
	like := &model.Like{
		Pid: pid,
		Uid: uid,
	}
	return dao.db.Where("pid = ? AND uid =?", pid, like).Delete(like).Error
}

func (dao *LikesDao) IsLiked(pid, uid int64) (bool, error) {

	err := dao.db.Where("pid = ? AND uid =?", pid, uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
