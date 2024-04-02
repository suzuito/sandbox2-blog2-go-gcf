package sandbox2blog2gogcf

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
)

func StartImageProcess(ctx context.Context, e event.Event) error {
	msg, err := initCloudEventFunction(ctx, &e)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	logger.Info("Hello world!", "msg", msg)
	return nil
}
