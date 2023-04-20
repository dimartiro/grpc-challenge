package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/dimartiro/grpc-examples/protos"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type ChallengeServer struct {
	pb.UnimplementedChallengeServer
}

func (s *ChallengeServer) Ping(ctx context.Context, in *pb.Empty) (*pb.ServiceReply, error) {
	//Implementar respuesta de ping y cambiar el primer nil por esa respuesta
	return nil, nil
}

func (s *ChallengeServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	//Implementar suma de num1 con num2 y cambiar el primer nil por esa respuesta
	//Recordar que al sumar un int32 + int32 el valor resultante puede dar overflow por lo que debemos usar int64
	//Para sumar 2 int64 que son int32 debemos castearlos haciendo: int64(valor1) + int64(valor2)
	return nil, nil
}

func (s *ChallengeServer) Forward(ctx context.Context, in *pb.ForwardRequest) (*pb.ServiceReply, error) {
	//NO MODIFICAR!
	//Crea la conexion con el servidor host especificado en el request donde vamos a ejecutar el forward
	conn, err := grpc.Dial(in.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewChallengeClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	//Respuesta del servicio
	var reply *pb.ServiceReply

	//Recibimos un forward, por lo que debemos evaluar que metodo ejecutar
	switch in.GetMethod() {
	case pb.Method_PING:
		//Si es un ping llamamos al ping en el servidor de destino
		reply, err = client.Ping(ctx, &pb.Empty{})
	case pb.Method_FORWARD:
		//Descomentar la siguiente linea agregando el campo que creaste en el mensaje Forward del proto
		//reply, err = client.Forward(ctx, in.CAMPO_QUE_DEBES_AGREGAR)
	default:
		err = errors.New("method not found")
	}

	if err != nil {
		log.Println("Error executing request", err)
		return nil, err
	}
	return reply, nil
}

// NO MODIFICAR
func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.CallerKey = zapcore.OmitKey
	cfg.EncoderConfig.MessageKey = zapcore.OmitKey
	cfg.EncoderConfig.TimeKey = zapcore.OmitKey

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger),
		),
	))

	pb.RegisterChallengeServer(server, &ChallengeServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
