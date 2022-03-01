package demo

import (
	"context"
	"fmt"

	"github.com/shinshin8/golang-grpc-middleware/demo/democtx"

	"google.golang.org/grpc"
)

// DemoServerInterceptor returns a new unary interceptor that put email address into a context.
func DemoServerInterceptor(d string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		v := democtx.GetDemo(ctx)

		var val string
		if v == "" {
			val = d
		} else {
			val = fmt.Sprintf("%s, %s", v, d)
		}

		ctx = democtx.SetDemo(ctx, val)

		return handler(ctx, req)
	}
}

func DemoErrorServerIntercepor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			fmt.Sprintf("発生したエラーです。： %s", err.Error())
		}
		return resp, err
	}
}
