// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(shutdown chan os.Signal, log *zap.SugaredLogger) *web.App {
	app := web.NewApp(shutdown)

	app.Handle(http.MethodGet, "/readiness", readiness)

	return app
}
