package grpcPlugin

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// CreateCred is helper to create creadential for grpc server
func CreateCred(certFile, keyFile string) (grpc.ServerOption, error) {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return grpc.ServerOption(grpc.Creds(creds)), nil
}
