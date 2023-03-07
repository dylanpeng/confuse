package grpc

import (
	"confuse/lib/proto/common"
	"confuse/lib/proto/common_service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"runtime"
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

	cfg2 = &Config{
		Host: "localhost",
		Port: 39572,
	}

	server  *Server
	server2 *Server
)

// 日志拦截器
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//t := time.Now()
	//fmt.Printf("gRpc begin method: method: %s | req: %v | time: %s", info.FullMethod, req, t.Format("2006-01-02 15:04:05.000000"))
	//fmt.Println()
	resp, err = handler(ctx, req)
	//fmt.Printf("gRpc finish method: %s | rsp: %v | time: %s | durations: %s", info.FullMethod, resp, t.Format("2006-01-02 15:04:05.000000"), time.Since(t))
	//fmt.Println()
	return
}

func init() {
	server = NewServer(cfg, &router{}, nil, grpc.ChainUnaryInterceptor(LoggerInterceptor))
	server2 = NewServer(cfg2, &router{}, nil, grpc.ChainUnaryInterceptor(LoggerInterceptor))

	for k, v := range clientConfig {
		vC, ok := clients[k]

		if ok {
			vC.Init(v)
		}
	}

	server.Start()
	server2.Start()
}

func TestServer(t *testing.T) {
	server.Start()
	server2.Start()

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

type IClient interface {
	Init(*ClientConfig)
	GetGrpcClient() (interface{}, *ClientConn, error)
}

type ClientConfig struct {
	Addrs    []string `toml:"addrs" json:"addrs"`
	Capacity int      `toml:"capacity" json:"capacity"`
	Idle     int      `toml:"idle" json:"idle"`
	Ttl      int      `toml:"ttl" json:"ttl"`
}

var (
	clients = map[string]IClient{
		"test": &GrpcHelloWorldClient{},
	}

	clientConfig = map[string]*ClientConfig{
		"test": &ClientConfig{
			Addrs:    []string{"localhost:39571", "localhost:39572"},
			Capacity: 1000,
			Idle:     300,
			Ttl:      300,
		},
	}
)

type GrpcHelloWorldClient struct {
	cfg   *ClientConfig
	pools []*Pool
}

func (g *GrpcHelloWorldClient) Init(c *ClientConfig) {
	g.cfg = c
	pools := make([]*Pool, 0, len(g.cfg.Addrs))

	for _, addr := range g.cfg.Addrs {
		pool := NewPool(DefaultFactory, addr, g.cfg.Capacity, time.Duration(g.cfg.Idle)*time.Second, time.Duration(g.cfg.Ttl)*time.Second)
		pools = append(pools, pool)
	}

	g.pools = pools
}

func (g *GrpcHelloWorldClient) GetGrpcClient() (interface{}, *ClientConn, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(g.pools))

	conn, err := g.pools[index].Get()

	if err != nil {
		return nil, nil, err
	}

	return common_service.NewCommonServiceClient(conn), conn, nil
}

func BenchmarkPool(b *testing.B) {
	client, ok := clients["test"]

	b.Logf("%d", runtime.NumCPU())
	if ok {
		// 通过 b.SetParallelism 方法实现对并发度的控制，例如执行b.SetParallelism(2) 则意味着并发度为 2*GOMAXPROCS
		b.SetParallelism(50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				gClient, conn, _ := client.GetGrpcClient()

				gC := gClient.(common_service.CommonServiceClient)

				req := &common.Empty{}
				rsp := &common.Response{}
				rsp, err := gC.CommonTest(context.Background(), req)

				if err != nil {
					b.Fatalf("call CommonTest fail. | err: %s", err)
				}

				b.Logf("call grpc common test. | rsp: %s | server: %s | id: %d", rsp.Message, conn.pool.addr, conn.Id)
				conn.Release()
			}
		})

	}
}
