package dao

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	"gorm.io/gorm"
	"sync"
)

type LikesDao struct {
	db *gorm.DB
}

type LikesDaoImpl interface {
	Insert(tags *model.Likes) error
	DelLikes(pid, uid int64) error
	CheckLikes(pid, uid int64) (bool, error)
	GetLikesCount(pid int64) int64
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

func (dao *LikesDao) Insert(likes *model.Likes) error {
	return dao.db.Create(likes).Error
}

func (dao *LikesDao) DelLikes(pid, uid int64) error {
	like := &model.Likes{
		Pid: pid,
		Uid: uid,
	}
	return dao.db.Where("pid = ? AND uid =?", pid, like).Delete(like).Error
}

func (dao *LikesDao) CheckLikes(pid, uid int64) (bool, error) {

	exist := dao.db.Where("pid = ? AND uid =?", pid, uid)
	if exist != nil {
		return false, fmt.Errorf("重复收藏")
	}
	return true, nil
}

func (dao *LikesDao) GetLikesCount(pid int64) int64 {
	var count int64
	dao.db.Model(&model.Likes{}).Where("pid = ?", pid).Count(&count)
	return count
}
