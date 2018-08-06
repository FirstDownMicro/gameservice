package main

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/Nanixel/FirstDownMicro/gameservice/proto"
	"github.com/micro/go-micro/broker"
)

const topic = "game.runplay"

type service struct {
	PubSub broker.Broker
}

func (s *service) StartGame(ctx context.Context, req *pb.Empty, res *pb.EmptyResponse) error {

	fmt.Println("Create called")

	return nil
}

func (s *service) RunPlay(ctx context.Context, req *pb.Empty, res *pb.EmptyResponse) error {
	fmt.Println("Run Play Called")

	// send broker message here
	s.publishEvent()

	return nil
}

func (s *service) publishEvent() error {

	samplejson := map[string]string{
		"name": "jesse",
	}

	sj, _ := json.Marshal(samplejson)

	msg := &broker.Message{
		Header: map[string]string{
			"id": "Test",
		},
		Body: sj,
	}

	if err := s.PubSub.Publish(topic, msg); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
