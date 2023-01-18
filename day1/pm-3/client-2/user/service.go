package user

import (
	"log"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/config"
	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client model.UsersClient

func init() {
	client = serviceUser()
}

func GetService() model.UsersClient {
	if client == nil {
		client = serviceUser()
	}

	return client
}

func serviceUser() model.UsersClient {
	log.Println("dialing service user")

	conn, err := grpc.Dial(config.SERVICE_USER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("fail to dial:", err)
	}

	return model.NewUsersClient(conn)
}
