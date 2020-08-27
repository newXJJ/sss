package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	GetImageCd "sss/GetImageCd/proto/GetImageCd"
)

type GetImageCd struct{}

func (e *GetImageCd) Handle(ctx context.Context, msg *GetImageCd.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GetImageCd.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
