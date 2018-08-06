package main

import (
	"fmt"

	pb "github.com/Nanixel/FirstDownMicro/gameservice/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
)

func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.gameservice"),
		micro.Version("latest"),
	)

	srv.Init()

	pubsub := srv.Server().Options().Broker

	pb.RegisterGameServiceHandler(srv.Server(), &service{PubSub: pubsub})

	fmt.Println("Gameservice starting on port :50051")

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
