package sandbox2blog2gogcf

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
)

func StartFileUploadedProcess(ctx context.Context, e event.Event) error {
	msg, err := initCloudEventFunction(ctx, &e)
	if err != nil {
		logger.Error("", "err", err)
		return err
	}
	logger.Info("Hello world!", "msg", msg)
	if err := u.StartFileUploadedProcessFromGCF(ctx, msg.Message.Data); err != nil {
		logger.Error("", "err", err)
		return err
	}
	return nil
}
