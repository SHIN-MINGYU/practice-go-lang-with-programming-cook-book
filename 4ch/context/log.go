package context

import (
	"context"

	"github.com/apex/log"
)

type key int

// logFields is key to use logging
const logFields key = 0

func getFields(ctx context.Context) *log.Fields{
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext function return data-filled item from context, after input context and item

func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
    fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx,logFields,fields)
	return ctx, e
}


// WithField function add log fields to context
func WithField(ctx context.Context,key string, value interface{}) context.Context {
	return WithFields(ctx,log.Fields{key : value})
}


// WithFields function add many log fields in context

func WithFields(ctx context.Context,fields log.Fields) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields(){
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields,f)
	return ctx
}
