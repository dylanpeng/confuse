package grpc

import (
	"confuse/lib/proto/common"
	"confuse/lib/proto/common_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

type commonServer struct {
	common_service.UnimplementedCommonServiceServer
}

func (s *commonServer) CommonTest(context.Context, *common.Empty) (*common.Response, error) {
	rsp := &common.Response{
		Code:    200,
		Message: "success",
	}

	return rsp, nil
}

type router struct{}

func (r *router) RegGrpcService(server *grpc.Server) {
	common_service.RegisterCommonServiceServer(server, &commonServer{})
}

var (
	cfg = &Config{
		Host: "localhost",
		Port: 39571,
	}

	server *Server
)

// 日志拦截器
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	t := time.Now()
	fmt.Printf("gRpc begin method: method: %s | req: %v | time: %s", info.FullMethod, req, t.Format("2006-01-02 15:04:05.000000"))
	fmt.Println()
	resp, err = handler(ctx, req)
	fmt.Printf("gRpc finish method: %s | rsp: %v | time: %s | durations: %s", info.FullMethod, resp, t.Format("2006-01-02 15:04:05.000000"), time.Since(t))
	fmt.Println()
	return
}

func init() {
	server = NewServer(cfg, &router{}, nil, grpc.ChainUnaryInterceptor(LoggerInterceptor))
}

func TestServer(t *testing.T) {
	server.Start()

	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:39571", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatalf("grpc dial fail. | err: %s", err)
	}

	client := common_service.NewCommonServiceClient(conn)

	rsp, err := client.CommonTest(context.Background(), &common.Empty{})

	if err != nil {
		t.Fatalf("grpc client call fail. | err: %s", err)
	}

	t.Logf("call CommonTest success. | rsp: %s", rsp.Message)
}
