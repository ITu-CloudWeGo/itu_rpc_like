package main

import (
	"context"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/dao"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	likes_service "github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service"
)

// LikesServiceImpl implements the last service interface defined in the IDL.
type LikesServiceImpl struct{}

//用户层面

// AddLikes implements the LikesServiceImpl interface.
func (s *LikesServiceImpl) AddLikes(ctx context.Context, req *likes_service.AddLikesRequest) (resp *likes_service.AddLikesResponse, err error) {
	// TODO: Your code here...

	LikesDao := dao.GetLikesDao()

	exist, err := LikesDao.CheckLikes(req.Pid, req.Uid)
	if exist == false {
		return nil, err
	}
	if err := LikesDao.Insert(&model.Likes{
		Pid: req.Pid,
		Uid: req.Uid,
	}); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &likes_service.AddLikesResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}

// DelLikes implements the LikesServiceImpl interface.
func (s *LikesServiceImpl) DelLikes(ctx context.Context, req *likes_service.DelLikesRequest) (resp *likes_service.DelLikesResponse, err error) {
	// TODO: Your code here...

	LikesDao := dao.GetLikesDao()

	if err := LikesDao.DelLikes(req.Pid, req.Uid); err != nil {
		return nil, fmt.Errorf("failed to delete likes: %w", err)
	}

	return &likes_service.DelLikesResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}

// 创建帖子的时候调用

// CreateLikes implements the LikesServiceImpl interface.
func (s *LikesServiceImpl) CreateLikes(ctx context.Context, req *likes_service.CreateLikesRequest) (resp *likes_service.CreateLikesResponse, err error) {
	// TODO: Your code here...
	PLikesDao := dao.GetLikesDao()
	err = PLikesDao.Insert(&model.Likes{
		Pid: req.Pid,
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	return &likes_service.CreateLikesResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}

//获取帖子点赞数,获取帖子的时候调用

// GetLikesCount implements the LikesServiceImpl interface.
func (s *LikesServiceImpl) GetLikesCount(ctx context.Context, req *likes_service.GetLikesCountRequest) (resp *likes_service.GetLikesCountResponse, err error) {
	// TODO: Your code here...
	var count int64
	PLikesDao := dao.GetLikesDao()
	count = PLikesDao.GetLikesCount(req.Pid)
	return &likes_service.GetLikesCountResponse{
		Status: 200,
		Msg:    "success",
		Count:  count,
	}, nil
}
