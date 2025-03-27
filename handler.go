package main

import (
	"context"
	"fmt"
	dao "github.com/ITu-CloudWeGo/itu_rpc_like/db"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/module"
	likes_service "github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service"
)

// LikesServiceImpl implements the last service interface defined in the IDL.
type LikesServiceImpl struct{}

// AddLikes implements the LikesServiceImpl interface.
func (s *LikesServiceImpl) AddLikes(ctx context.Context, req *likes_service.AddLikesRequest) (resp *likes_service.AddLikesResponse, err error) {
	// TODO: Your code here...
	LikesDao := dao.GetLikesDaoImpl()
	exist, err := LikesDao.CheckLikes(req.Pid, req.Uid)
	if exist == false {
		return nil, err
	}
	if err := LikesDao.Insert(&module.Likes{
		Pid: req.Pid,
		Uid: req.Uid,
	}); err != nil {
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
	LikesDao := dao.GetLikesDaoImpl()
	if err := LikesDao.DelLikes(req.Pid, req.Uid); err != nil {
		return nil, fmt.Errorf("failed to delete likes: %w", err)
	}

	return &likes_service.DelLikesResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
