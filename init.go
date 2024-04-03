package sandbox2blog2gogcf

import (
	"context"
	"log/slog"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/pkg/environment"
	"github.com/suzuito/sandbox2-go/blog2/pkg/inject"
	"github.com/suzuito/sandbox2-go/blog2/pkg/usecase"
)

var u usecase.Usecase
var logger *slog.Logger = slog.Default()

func init() {
	ctx := context.Background()
	var err error
	env := environment.Environment{}
	if err := envconfig.Process("", &env); err != nil {
		logger.Error("", "err", err)
		return
	}
	u, logger, err = inject.NewUsecaseImpl(ctx, &env)
	if err != nil {
		logger.Error("", "err", err)
		return
	}
}
