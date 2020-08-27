package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	GA "sss/GetArea/proto/GetArea"
)

type GetArea struct{}

func (e *GetArea) Handle(ctx context.Context, msg *GA.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GA.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
