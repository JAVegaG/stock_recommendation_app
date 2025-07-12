package utils

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/httplog/v3"
	"github.com/go-chi/traceid"
	"github.com/golang-cz/devslog"
)

func getEnv(envStr string) string {

	envStr = strings.ToLower(envStr)

	switch {
	case strings.HasPrefix("prod", envStr):
		return "prod"
	case strings.HasPrefix("test", envStr):
		return "test"
	case strings.HasPrefix("dev", envStr):
		return "dev"
	case strings.HasPrefix("local", envStr):
		return "local"
	default:
		return "dev"
	}
}

var DEPLOY_ENV = getEnv(os.Getenv("ENV"))
var IS_LOCALHOST = DEPLOY_ENV == "local"

var LOG_FORMAT = httplog.SchemaECS.Concise(IS_LOCALHOST)

func logHandler(isLocalhost bool, handlerOpts *slog.HandlerOptions) slog.Handler {
	if isLocalhost {
		// Pretty logs for localhost development.
		return devslog.NewHandler(os.Stdout, &devslog.Options{
			SortKeys:           true,
			MaxErrorStackTrace: 5,
			MaxSlicePrintSize:  20,
			HandlerOptions:     handlerOpts,
		})
	}

	// JSON logs for production with "traceId".
	return traceid.LogHandler(
		slog.NewJSONHandler(os.Stdout, handlerOpts),
	)
}

func getLogger() *slog.Logger {

	logger := slog.New(logHandler(IS_LOCALHOST, &slog.HandlerOptions{
		AddSource:   !IS_LOCALHOST,
		ReplaceAttr: LOG_FORMAT.ReplaceAttr,
	}))

	if !IS_LOCALHOST {
		logger = logger.With(
			slog.String("app", "Stocl Recommendation API"),
			slog.String("version", "v1.0.0"),
			slog.String("env", DEPLOY_ENV),
		)
	}

	return logger
}

var Logger = getLogger()

func getRequestLogger() func(http.Handler) http.Handler {
	return httplog.RequestLogger(Logger, &httplog.Options{
		// Level defines the verbosity of the request logs:
		// slog.LevelDebug - log all responses (incl. OPTIONS)
		// slog.LevelInfo  - log all responses (excl. OPTIONS)
		// slog.LevelWarn  - log 4xx and 5xx responses only (except for 429)
		// slog.LevelError - log 5xx responses only
		Level: slog.LevelInfo,

		// Log attributes using given schema/format.
		Schema: LOG_FORMAT,

		// RecoverPanics recovers from panics occurring in the underlying HTTP handlers
		// and middlewares. It returns HTTP 500 unless response status was already set.
		//
		// NOTE: Panics are logged as errors automatically, regardless of this setting.
		RecoverPanics: true,

		// Filter out some request logs.
		Skip: func(req *http.Request, respStatus int) bool {
			return respStatus == 404 || respStatus == 405
		},

		// Select request/response headers to be logged explicitly.
		LogRequestHeaders:  []string{"Origin"},
		LogResponseHeaders: []string{},

		// Log all requests with invalid payload as curl command.
		LogExtraAttrs: func(req *http.Request, reqBody string, respStatus int) []slog.Attr {
			if respStatus == 400 || respStatus == 422 {
				req.Header.Del("Authorization")
				return []slog.Attr{slog.String("curl", httplog.CURL(req, reqBody))}
			}
			return nil
		},
	})
}

var HttpRequestLogger = getRequestLogger()
