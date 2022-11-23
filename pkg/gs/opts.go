package gs

import (
	"context"
	"github.com/Charliego93/go-i18n/v2"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/transerver/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func WithI18nOpts(opts ...i18n.Option) Option {
	return func(gs *Server) {
		gs.i18nOpts = opts
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

// Interceptor

// UnaryI18nInterceptor add accept-language key with language.Tag to context
func UnaryI18nInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = context.WithValue(ctx, "accept-language", utils.Language(ctx))
	return handler(ctx, req)
}

// StreamI18nInterceptor add accept-language key with language.Tag to context
func StreamI18nInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = ss.SetHeader(metadata.Pairs("accept-language", utils.Language(ss.Context()).String()))
	return handler(srv, ss)
}
