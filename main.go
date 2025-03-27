package main

import (
	"github.com/ITu-CloudWeGo/itu_rpc_like/config"
	likes_service "github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service/likesservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {
	conf := config.GetConfig()

	r, err := etcd.NewEtcdRegistry(
		[]string{conf.Registry.RegistryAddress},
		etcd.WithDialTimeoutOpt(10*time.Second),
	)
	if err != nil {
		log.Fatalf("初始化 Etcd 注册中心失败: %v", err)
	}

	svr := likes_service.NewServer(
		new(LikesServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "likes",
		}),
	)

	if err := svr.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
