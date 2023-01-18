package garage

import (
	"log"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/config"
	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client model.GaragesClient

func init() {
	client = serviceGarage()
}

func GetService() model.GaragesClient {
	if client == nil {
		client = serviceGarage()
	}

	return client
}

func serviceGarage() model.GaragesClient {
	log.Println("dialing service garage")

	conn, err := grpc.Dial(config.SERVICE_GARAGE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail to dial", err)
	}

	return model.NewGaragesClient(conn)
}
