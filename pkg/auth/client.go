package auth

import (
	"api-gateway/configs"
	"api-gateway/pkg/auth/pb"
	"fmt"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *configs.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
