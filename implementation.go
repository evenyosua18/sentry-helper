package sentry_helper

import (
	"context"
	"github.com/getsentry/sentry-go"
)

func (h *Helper) LogError(span interface{}, err error, status ...int) {
	sp := span.(*sentry.Span)

	stat := 0

	if len(status) != 0 {
		stat = status[0]
	}

	LogError(sp, err, stat)
}

func (h *Helper) LogObject(span interface{}, name string, obj interface{}) {
	sp := span.(*sentry.Span)
	LogObject(sp, name, obj)
}

func (h *Helper) StartParent(ctx interface{}) interface{} {
	return StartParent(ctx)
}

func (h *Helper) StartChild(ctx context.Context, request ...interface{}) interface{} {
	return StartChild(ctx, request)
}

func (h *Helper) Close(span interface{}) {
	sp := span.(*sentry.Span)
	sp.Finish()
}

func (h *Helper) Context(span interface{}) context.Context {
	sp := span.(*sentry.Span)
	return sp.Context()
}

func (h *Helper) GetTraceID(span interface{}) string {
	sp := span.(*sentry.Span)
	return sp.TraceID.String()
}

//	Close(span interface{})
//	Context() context.Context
