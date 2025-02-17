package main

import (
	"context"
	"fmt"
	"net"

	"github.com/TremblingV5/DouTok/applications/relation/handler"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	service.InitViper()
	service.InitRedisClient()
	service.InitSyncProducer()
	service.InitConsumerGroup()
	service.InitId()
	service.InitDB()
	service.InitSafeMap()
	service.InitMutex()
	rpc.Init()
	//go service.Flush()
}

func main() {
	Init()
	defer func() {
		_ = service.SyncProducer.Close()
	}()

	// 启动 kafka 消费者协程，消费点赞消息
	go service.ConsumeMsg()

	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)

	ServiceName := service.ViperConfig.Viper.GetString("Server.Name")
	ServiceAddr := fmt.Sprintf("%s:%d", service.ViperConfig.Viper.GetString("Server.Address"), service.ViperConfig.Viper.GetInt("Server.Port"))
	EtcdAddress := fmt.Sprintf("%s:%d", service.ViperConfig.Viper.GetString("Etcd.Address"), service.ViperConfig.Viper.GetInt("Etcd.Port"))

	r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(ServiceName),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%s", service.ViperConfig.Viper.GetString("Otel.Host"), service.ViperConfig.Viper.GetString("Otel.Port"))),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	svr := relationservice.NewServer(
		new(handler.RelationServiceImpl),
		server.WithServiceAddr(addr),                                       // address
		server.WithMiddleware(middleware.CommonMiddleware),                 // middleware
		server.WithMiddleware(middleware.ServerMiddleware),                 // middleware
		server.WithRegistry(r),                                             // registry
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(tracing.NewServerSuite()),                         // trace
		// Please keep the same as provider.WithServiceName
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)

	if err := svr.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}
