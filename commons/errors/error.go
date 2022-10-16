package errors

import (
	"context"
	"github.com/Charliego93/go-i18n/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	Code     codes.Code
	Template string
	Data     any
}

func (e Err) Error() string {
	return e.Template
}

func (e Err) Localize(ctx context.Context) string {
	return i18n.MustTr(ctx, &i18n.LocalizeConfig{
		MessageID:    e.Template,
		TemplateData: e.Data,
	})
}

type Option func(*Err)

func WithCode(code codes.Code) Option {
	return func(e *Err) {
		e.Code = code
	}
}

func WithData(data any) Option {
	return func(e *Err) {
		e.Data = data
	}
}

func WithTemplate(template string) Option {
	return func(e *Err) {
		e.Template = template
	}
}

func New(template string, opts ...Option) error {
	e := &Err{Template: template, Code: codes.Internal}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func ErrorArgument(ctx context.Context, err error) error {
	if e, ok := err.(interface{ Field() string }); ok {
		return ErrorArgumentf(ctx, &i18n.LocalizeConfig{
			MessageID:    "InvalidArgument",
			TemplateData: e.Field(),
		})
	}
	return err
}

func ErrorArgumentf[T i18n.MessageID](ctx context.Context, messageId T) error {
	return status.Error(codes.InvalidArgument, i18n.MustTr(ctx, messageId))
}
