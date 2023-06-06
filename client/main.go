package main

import (
	"context"
	"fmt"
	pb "go_study/grpc_study/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//type PerRPCCredentials interface {
//	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
//	RequireTransportSecurity() bool
//}

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "id",
		"appKey": "key",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// TLS认证, 自签名证书文件和访问域名(配合域名校验)
	//creds, _ := credentials.NewClientTLSFromFile("D:\\GoPath\\src\\go_study\\grpc_study\\key\\test.pem", "*.grpcstudy.com")

	// 连接server端, 此处禁用安全传输, 即没有加密和验证
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))

	// Token认证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:9090", opts...)

	if err != nil {
		log.Fatalf("connect failed: #{err}")
	}

	// 关闭连接
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行RPC调用, 该方法在服务端实现并返回结果
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "GRPC"})

	fmt.Println(resp.GetResponseMsg())
}
