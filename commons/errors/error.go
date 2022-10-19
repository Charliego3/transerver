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
	return i18n.MustTr(ctx, &i18n.Localized{
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

func New[T i18n.Message](ctx context.Context, code codes.Code, messageId T) error {
	return status.Error(code, i18n.MustTr(ctx, messageId))
}

func NewInternal[T i18n.Message](ctx context.Context, messageId T) error {
	return New(ctx, codes.Internal, messageId)
}

func NewArgument(ctx context.Context, err error) error {
	if e, ok := err.(interface{ Field() string }); ok {
		return NewArgumentf(ctx, &i18n.Localized{
			MessageID:    "InvalidArgument",
			TemplateData: e.Field(),
		})
	}
	return err
}

func NewArgumentf[T i18n.Message](ctx context.Context, messageId T, data ...any) error {
	if len(data) > 0 {
		if id, ok := ((any)(messageId)).(string); ok {
			return New(ctx, codes.InvalidArgument, &i18n.Localized{
				MessageID:    id,
				TemplateData: data[0],
			})
		}
	}
	return New(ctx, codes.InvalidArgument, messageId)
}
