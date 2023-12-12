package grpc

import (
	ssov1 "github.com/irsCooper/gRPC-postgresql-go/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type Notepad interface {
	Set(ownerID, title, content string, tags ...string) (articleID int64, err error)
	One(articleID int64) (bool, error)
	All(articleID int64) (bool, error)
	IsOwner(ownerID int64) (bool, error)
	Delete(articleID int64) (string, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	notepad Notepad
}

func Register(gRPC *grpc.Server, notepad Notepad) {
	ssov1.RegisterAuthServer(gRPC, serverAPI{notepad: notepad})
}

const (
	emptyValue = 0
)

