package engine

import (
	"github.com/runabol/tork/pkg/middleware/job"
	"github.com/runabol/tork/pkg/middleware/node"
	"github.com/runabol/tork/pkg/middleware/task"
	"github.com/runabol/tork/pkg/middleware/web"
)

var defaultEngine *Engine = New(Config{})

func RegisterWebMiddleware(mw web.MiddlewareFunc) {
	defaultEngine.RegisterWebMiddleware(mw)
}

func RegisterTaskMiddleware(mw task.MiddlewareFunc) {
	defaultEngine.RegisterTaskMiddleware(mw)
}

func RegisterJobMiddleware(mw job.MiddlewareFunc) {
	defaultEngine.RegisterJobMiddleware(mw)
}

func RegisterNodeMiddleware(mw node.MiddlewareFunc) {
	defaultEngine.RegisterNodeMiddleware(mw)
}

func RegisterEndpoint(method, path string, handler web.HandlerFunc) {
	defaultEngine.RegisterEndpoint(method, path, handler)
}

func Start() error {
	return defaultEngine.Start()
}

func Terminate() error {
	return defaultEngine.Terminate()
}

func SetMode(mode Mode) {
	defaultEngine.SetMode(mode)
}

func Run() error {
	return defaultEngine.Run()
}