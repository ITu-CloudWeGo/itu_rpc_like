package dao

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_like/config"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/module"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type LikesDao struct {
	db *gorm.DB
}

type LikesDaoImpl interface {
	Insert(tags *module.Likes) error
	DelLikes(pid, uid int64) error
	CheckLikes(pid, uid int64) (bool, error)
}

var (
	instanceLikesDAO *LikesDao
	onceLikesDao     sync.Once
)

func GetLikesDao() LikesDaoImpl {
	onceLikesDao.Do(func() {
		instanceLikesDAO = createDB()
	})
	return instanceLikesDAO
}
func createDB() *LikesDao {
	conf := config.Config{}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.PostgresSQL.Host,
		conf.PostgresSQL.User,
		conf.PostgresSQL.Password,
		conf.PostgresSQL.DBName,
		conf.PostgresSQL.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	err = db.AutoMigrate(&module.Likes{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	instanceLikesDAO = &LikesDao{
		db: db,
	}
	return instanceLikesDAO
}
func (dao *LikesDao) Insert(tags *module.Likes) error {
	return dao.db.Create(tags).Error
}

func (dao *LikesDao) DelLikes(pid, uid int64) error {
	tags := &module.Likes{
		Pid: pid,
		Uid: uid,
	}
	return dao.db.Where("pid = ? AND uid =?", pid, tags).Delete(tags).Error
}

func (dao *LikesDao) CheckLikes(pid, uid int64) (bool, error) {

	exist := dao.db.Where("pid = ? AND uid =?", pid, uid)
	if exist != nil {
		return false, fmt.Errorf("重复收藏")
	}
	return true, nil
}
