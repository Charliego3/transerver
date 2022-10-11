package gs

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
)

type Option func(gs *Server)

func WithServerOption(opts ...grpc.ServerOption) Option {
	return func(gs *Server) {
		gs.serverOpts = append(opts, gs.serverOpts...)
	}
}

func WithAuthFunc(fn grpc_auth.AuthFunc) Option {
	return func(gs *Server) {
		gs.authFunc = fn
	}
}

func WithCtxTags(opts ...grpc_ctxtags.Option) Option {
	return func(gs *Server) {
		gs.ctxTagOpts = opts
	}
}

func WithTracing(opts ...grpc_opentracing.Option) Option {
	return func(gs *Server) {
		gs.tracingOpt = opts
	}
}

func WithLoggerOpt(opts ...grpc_zap.Option) Option {
	return func(gs *Server) {
		gs.loggerOpts = opts
	}
}

func WithRecoveryHandlerFuncContext(fn grpc_recovery.RecoveryHandlerFuncContext) Option {
	return func(gs *Server) {
		gs.recoverOpt = []grpc_recovery.Option{grpc_recovery.WithRecoveryHandlerContext(fn)}
	}
}

func WithStreamServerInterceptor(interceptor grpc.StreamServerInterceptor) Option {
	return func(gs *Server) {
		gs.streamOpts = append(gs.streamOpts, interceptor)
	}
}

func WithUnaryServerInterceptor(interceptor grpc.UnaryServerInterceptor) Option {
	return func(gs *Server) {
		gs.unaryOpts = append(gs.unaryOpts, interceptor)
	}
}
