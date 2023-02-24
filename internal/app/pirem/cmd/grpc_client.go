package cmd

import (
	"context"
	"net"
	"time"

	apiremv1 "github.com/NaKa2355/pirem/pkg/apirem.v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func MakeConnection(protocol string, addr string) (*grpc.ClientConn, apiremv1.PiRemServiceClient, error) {
	var conn *grpc.ClientConn
	var client apiremv1.PiRemServiceClient
	credentials := insecure.NewCredentials()

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, protocol, addr)
	}

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials),
		grpc.WithBlock(),
		grpc.WithContextDialer(dialer),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, DomainSocketPath, options...)
	if err != nil {
		return conn, client, err
	}
	client = apiremv1.NewPiRemServiceClient(conn)
	return conn, client, nil
}

func MarshalToString(m protoreflect.ProtoMessage) (string, error) {
	op := protojson.MarshalOptions{Indent: "  "}
	raw, err := op.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(raw), err
}
