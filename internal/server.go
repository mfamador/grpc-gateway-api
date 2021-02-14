package echo

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mfamador/api/v1/internal/gen"
	"google.golang.org/grpc"
	"net"
	"net/http"
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
	go grpcServer.Serve(grpcLis)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = echo.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

//
//func Run() error {
//	EndPoint = ":" + ServerPort
//	conn, err := net.Listen("tcp", EndPoint)
//	if err != nil {
//		log.Printf("TCP Listen err:%v\n", err)
//	}
//	srv := newServer(conn)
//
//	log.Info().Str("gRPC and http listen on: %s\n", ServerPort)
//	if err = srv.Serve(conn); err != nil {
//		log.Printf("Serve: %v\n", err)
//	}
//
//	return err
//}
//
//func newServer(conn net.Listener) *http.Server {
//	grpcServer := newGrpc()
//	gwmux, err := newGateway()
//	if err != nil {
//		panic(err)
//	}
//	mux := http.NewServeMux()
//	mux.Handle("/", gwmux)
//	//mux.HandleFunc("/swagger/", serveSwaggerFile)
//	//serveSwaggerUI(mux)
//
//	return &http.Server{
//		Addr:    EndPoint,
//		Handler: GrpcHandlerFunc(grpcServer, mux),
//	}
//}
//
//func newGrpc() *grpc.Server {
//	server := grpc.NewServer()
//	echo.RegisterEchoServiceServer(server, NewEchoService())
//	return server
//}
//
//func newGateway() (http.Handler, error) {
//	ctx := context.Background()
//	gwmux := runtime.NewServeMux()
//	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
//	if err := echo.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, EndPoint, dialOpts); err != nil {
//		return nil, err
//	}
//	return gwmux, nil
//}
//
//func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
//	log.Print("GrpcHandlerFunc")
//	if otherHandler == nil {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			grpcServer.ServeHTTP(w, r)
//		})
//	}
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
//			grpcServer.ServeHTTP(w, r)
//		} else {
//			otherHandler.ServeHTTP(w, r)
//		}
//	})
//}
