package log

import "context"

type key int

var k key

func NewContext(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, k, val)
}

func FromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(k).(string)
	return val, ok
}
