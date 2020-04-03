package api
import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "Details"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DetailsClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewDetailsClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
