package echo

import (
	"context"
	"flag"
	echo "github.com/mfamador/api/v1/internal/gen"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	ServerPort         string
	SwaggerDir         string
	EndPoint           = "8181"
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer()
	echoService := NewEchoService()
	echo.RegisterEchoServiceServer(grpcServer, echoService)

	grpcLis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return err
	}
	go func() {
		err := grpcServer.Serve(grpcLis)
		if err != nil {
			log.Err(err)
		}
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = echo.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
