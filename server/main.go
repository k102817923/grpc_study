package main

import (
	"context"
	"errors"
	"fmt"
	pb "go_study/grpc_study/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("未传输Token")
	}

	var appId string
	var appKey string

	if v, ok := md["appid"]; ok {
		appId = v[0]
	}

	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	if appId != "id" || appKey != "key" {
		return nil, errors.New("Token错误")
	}

	fmt.Printf("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// TLS认证, 自签名证书文件和私钥文件
	//creds, _ := credentials.NewServerTLSFromFile("D:\\GoPath\\src\\go_study\\grpc_study\\key\\test.pem", "D:\\GoPath\\src\\go_study\\grpc_study\\key\\test.key")

	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建GRPC服务
	//grpcServer := grpc.NewServer()
	//grpcServer := grpc.NewServer(grpc.Creds(creds))

	// Token认证
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	// 在GRPC服务端中注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err := grpcServer.Serve(listen)

	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
}
