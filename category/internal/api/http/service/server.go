package service

import (
	"context"
	"github.com/buaazp/fasthttprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
	"net/http/pprof"
)

const pprofUrlPrefix = "/debug/pprof"

type Server struct {
	logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) ListenAndServe(ctx context.Context, listen string, enablePprof bool) error {
	router := fasthttprouter.New()

	router.GET("/healthz", s.HandleLivenessProbe)
	router.GET("/readyz", s.HandleReadinessProbe)
	router.GET("/metrics", s.HandleMetrics)
	// обязательный путь для проверки валидной работы самого HTTP-сервера
	router.GET("/ping", s.HandlePing)

	if enablePprof {
		for _, path := range []string{"/", "/allocs", "/block", "/goroutine", "/heap", "/mutex", "/threadcreate"} {
			router.GET(pprofUrlPrefix+path, fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
		}
		router.GET(pprofUrlPrefix+"/cmdline", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline))
		router.GET(pprofUrlPrefix+"/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
		router.GET(pprofUrlPrefix+"/symbol", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol))
		router.GET(pprofUrlPrefix+"/trace", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace))
		s.logger.Warn("pprof routes registered", zap.String("name", pprofUrlPrefix))
	}

	server := &fasthttp.Server{
		Handler: router.Handler,
	}

	go func() {
		<-ctx.Done()
		err := server.Shutdown()
		if err != nil {
			s.logger.Error("server shutdown with error", zap.Error(err))
		}
	}()
	return server.ListenAndServe(listen)
}

func (s *Server) HandleLivenessProbe(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/plain; charset=utf-8", "OK")
}

func (s *Server) HandleReadinessProbe(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/plain; charset=utf-8", "OK")
}

func (s *Server) HandleMetrics(ctx *fasthttp.RequestCtx) {
	prometheusHandler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	prometheusHandler(ctx)
}

func (s *Server) HandlePing(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/plain; charset=utf-8", "PONG")
}
