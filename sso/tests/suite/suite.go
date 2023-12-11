package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	ssov1 "github.com/irsCooper/gRPC-postgresql-go/protos/gen/go/sso"
	"github.com/irsCooper/gRPC-postgresql-go/sso/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T                  // from call methods
	Cfg        *config.Config   // configuration
	AuthClient ssov1.AuthClient // client for interaction from gRPC-server
}

const (
	grpcHost = "localhost"
)

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("./sso/config/local.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(
		context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)	
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T: t,
		Cfg: cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
